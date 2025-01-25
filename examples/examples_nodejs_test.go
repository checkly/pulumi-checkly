//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestBasicTs(t *testing.T) {
	opts := getJSBaseOptions(t).With(integration.ProgramTestOptions{
		Dir:        filepath.Join(getCwd(t), "ts"),
		NoParallel: true, // Otherwise resource IDs will conflict.
	})

	integration.ProgramTest(t, &opts)
}

func TestBasicJs(t *testing.T) {
	opts := getJSBaseOptions(t).With(integration.ProgramTestOptions{
		Dir:        filepath.Join(getCwd(t), "js"),
		NoParallel: true, // Otherwise resource IDs will conflict.
	})

	integration.ProgramTest(t, &opts)
}
