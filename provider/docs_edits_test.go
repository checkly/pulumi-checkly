package checkly

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// upstreamModule is the Go module of the bridged Terraform provider.
const upstreamModule = "github.com/checkly/terraform-provider-checkly/v2"

// upstreamDocs returns the upstream resource and data source docs of the
// terraform-provider-checkly version that go.mod pins, read from the Go module cache, keyed
// by file name.
func upstreamDocs(t *testing.T) map[string][]byte {
	t.Helper()
	out, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", upstreamModule).Output()
	if err != nil {
		t.Fatalf("locating %s: %v", upstreamModule, err)
	}
	dir := strings.TrimSpace(string(out))
	if dir == "" {
		t.Fatalf("%s is not in the module cache; run go mod download", upstreamModule)
	}
	docs := map[string][]byte{}
	for _, kind := range []string{"resources", "data-sources"} {
		files, err := filepath.Glob(filepath.Join(dir, "docs", kind, "*.md"))
		if err != nil {
			t.Fatal(err)
		}
		for _, f := range files {
			content, err := os.ReadFile(f)
			if err != nil {
				t.Fatal(err)
			}
			docs[filepath.Base(f)] = content
		}
	}
	if len(docs) == 0 {
		t.Fatalf("no upstream docs in %s", dir)
	}
	// checkly_maintenance_windows is not mapped, so its page is not used.
	delete(docs, "maintenance_windows.md")
	return docs
}

// terraformOnlyTerms are phrases that describe Terraform rather than Pulumi. None may remain
// in the prose of the edited docs, or in the comments of their examples, which the bridge
// converts into the Pulumi examples.
var terraformOnlyTerms = []string{
	"moved {", "moved block", "`moved`",
	"terraform apply", "plan output", "sensitive(", "sensitive = true", "sensitive variable",
	"replace_triggered_by", "filebase64", "first plan", "next apply", "every plan",
	"plan time", "apply time", "Terraform state", "Terraform 1.", "v2 upgrade guide",
	"Migrating an existing", "Terraform",
}

// allowedTerraformTerms are uses of terraformOnlyTerms that are correct for Pulumi users.
var allowedTerraformTerms = []string{
	// The upstream guide, named as the Terraform provider's.
	"the Terraform provider's v2 upgrade guide",
}

// ignoreChanges matches ignore_changes, which is correct only when named as the Pulumi
// resource option, not used as Terraform syntax.
var ignoreChanges = regexp.MustCompile("`ignore_changes`(?: resource option)?")

var (
	markdownLink = regexp.MustCompile(`\]\([^)]*\)`)
	codeFence    = regexp.MustCompile("(?s)```[a-z]*\\n(.*?)```")
	frontMatter  = regexp.MustCompile(`(?s)\A---\n.*?\n---\n`)
)

// proseAndComments returns the text of a doc outside code blocks and link targets, followed by
// the comment lines inside code blocks. Code is left out: the bridge converts it to Pulumi.
func proseAndComments(content []byte) string {
	// The bridge reads the page body only, so the front matter, which repeats the
	// description in plain text, is left out.
	content = frontMatter.ReplaceAll(content, nil)
	var comments []string
	prose := codeFence.ReplaceAllStringFunc(string(content), func(block string) string {
		for _, line := range strings.Split(block, "\n") {
			if trimmed := strings.TrimSpace(line); strings.HasPrefix(trimmed, "#") {
				comments = append(comments, trimmed)
			}
		}
		return ""
	})
	return markdownLink.ReplaceAllString(prose, "]") + "\n" + strings.Join(comments, "\n")
}

// TestDocsEditRulesMatchUpstream checks that every docs edit rule still matches the pinned
// upstream docs, and that no Terraform-only wording remains once they are applied. When
// upstream rewords a phrase, a rule stops matching and the Terraform wording would silently
// return to the generated docs; this test fails instead, so the rule can be updated along
// with the upstream version.
func TestDocsEditRulesMatchUpstream(t *testing.T) {
	docs := upstreamDocs(t)

	anyDoc := func(match func([]byte) bool) bool {
		for _, content := range docs {
			if match(content) {
				return true
			}
		}
		return false
	}
	for _, r := range docsReplacements {
		if !anyDoc(func(c []byte) bool { return bytes.Contains(c, []byte(r[0])) }) {
			t.Errorf("no upstream doc contains %q; update the rule in docs_edits.go", r[0])
		}
	}
	for _, r := range docsPatternReplacements {
		if !anyDoc(r.pattern.Match) {
			t.Errorf("no upstream doc matches %s; update the rule in docs_edits.go", r.pattern)
		}
	}
	for file := range movingSections {
		content, ok := docs[file]
		if !ok {
			t.Errorf("upstream has no %s; update movingSections in docs_edits.go", file)
			continue
		}
		if !movingSection.Match(content) {
			t.Errorf("%s has no \"## Moving to\" section before the generated schema; update movingSections", file)
		}
	}

	// Apply the rules as tfgen does, and check what remains in every doc.
	edits := docsEditRules(nil)
	for file, content := range docs {
		for _, e := range edits {
			if ok, _ := filepath.Match(e.Path, file); !ok {
				continue
			}
			var err error
			if content, err = e.Edit(file, content); err != nil {
				t.Fatalf("%s: %v", file, err)
			}
		}

		if bytes.Contains(content, []byte("moved {")) {
			t.Errorf("%s still has a moved block after the edits", file)
		}
		text := proseAndComments(content)
		for _, allowed := range allowedTerraformTerms {
			text = strings.ReplaceAll(text, allowed, "")
		}
		for _, term := range terraformOnlyTerms {
			for i := 0; ; {
				j := strings.Index(text[i:], term)
				if j < 0 {
					break
				}
				t.Errorf("%s still has %q after the edits: %q", file, term, excerpt(text, i+j))
				i += j + len(term)
			}
		}
		for _, m := range ignoreChanges.FindAllStringIndex(text, -1) {
			if !strings.HasSuffix(text[m[0]:m[1]], " resource option") {
				t.Errorf("%s uses ignore_changes as Terraform syntax: %q", file, excerpt(text, m[0]))
			}
		}
	}
}

func excerpt(text string, i int) string {
	return text[max(0, i-60):min(len(text), i+80)]
}
