# Contributing to Checkly Pulumi Provider

> Learn about our [Commitment to Open Source](https://checklyhq.com/oss).

Hi! We are really excited that you are interested in contributing to Checkly Pulumi Provider, and we really appreciate your commitment. Before submitting your contribution, please make sure to take a moment and read through the following guidelines:

- [Code of Conduct](./CODE-OF-CONDUCT.md)
- [Issue Reporting Guidelines](#issue-reporting-guidelines)
- [Pull Request Guidelines](#pull-request-guidelines)
- [Setting up your development environment](#setting-up-your-development-environment)
- [Committing Generated Code](#committing-generated-code)
- [Running Integration Tests](#running-integration-tests)

## Issue Reporting Guidelines

- Always use [GitHub issues](https://github.com/checkly/pulumi-checkly/issues/new/choose) to create new issues and select the corresponding issue template.

## Pull Request Guidelines

- Checkout a topic branch from a base branch, e.g. `main`, and merge back against that branch.

- [Make sure to tick the "Allow edits from maintainers" box](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/allowing-changes-to-a-pull-request-branch-created-from-a-fork). This allows us to directly make minor edits / refactors and saves a lot of time.

- Add accompanying documentation, usage samples & test cases
- Add/update demo files to showcase your changes.
- Use existing resources as templates and ensure that each property has a corresponding `description` field.
- Each PR should be linked with an issue, use [GitHub keywords](https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/using-keywords-in-issues-and-pull-requests) for that.
- Be sure to follow up project code styles

- If adding a new feature:
  - Provide a convincing reason to add this feature. Ideally, you should open a "feature request" issue first and have it approved before working on it (it should has the label "state: confirmed")
  - Each new feature should be linked to

- If fixing a bug:
  - Provide a detailed description of the bug in the PR. A working demo would be great!

- It's OK to have multiple small commits as you work on the PR - GitHub can automatically squash them before merging.

- Make sure tests pass!

- Commit messages must follow the [semantic commit messages](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716) so that changelogs can be automatically generated.

## Pull request review: Ship / Show / Ask

This repo supports the [Ship / Show / Ask](https://martinfowler.com/articles/ship-show-ask.html)
strategy. Signal your intent by ending the **PR title** with a tag (case-insensitive), e.g.
`feat: add retries [ship]`:

- **`[ship]`** — No review needed. `github-actions[bot]` auto-approves the PR; merge it yourself
  once CI is green. Use for low-risk changes you're confident in.
- **`[show]`** — Merge now, review after the fact. Also auto-approved; open it for visibility and
  merge once CI passes, inviting comments for follow-up.
- **`[ask]`** or **no tag** — Normal review. The PR waits for a human approval before it can merge.

Notes:

- **Auto-approval is limited to PRs opened from a branch in this repository** (i.e. members and
  collaborators with push access). Pull requests from forks are **never** auto-approved and always
  require manual review, regardless of the title tag.
- The tag only counts at the very end of the title. A `[ship]`/`[show]` elsewhere in the title is
  ignored — so e.g. a GitHub-generated revert title (`Revert "feat: x [ship]"`) still gets manual
  review.
- If you tag a PR `[ship]`/`[show]` and later change the title to `[ask]` (or drop the tag), the
  earlier auto-approval is dismissed so the PR returns to manual review.
- Dismissing the bot's approval by hand doesn't stick: while the title still ends with
  `[ship]`/`[show]`, the next push (or other PR update) re-approves. To durably demand review on a
  tagged PR, submit a **Request changes** review or retitle the PR to `[ask]`.
- A draft PR tagged `[ship]`/`[show]` is approved only once it's marked ready for review.
- This only adds an approving review; merging still respects branch protection (passing CI, etc.).
  It does not auto-merge.

The behavior lives in
[`.github/workflows/checkly-ship-show-ask.yml`](.github/workflows/checkly-ship-show-ask.yml); the
`checkly-` prefix keeps `make ci-mgmt` from deleting it along with other non-generated workflows.
It depends on two GitHub settings outside this repo: **Allow GitHub Actions to create and approve
pull requests** must be enabled, and branch protection on `main` must require an approving review
(a review count — not a Code Owners review, which a bot can't satisfy).

## Setting up your development environment

### Pulumi prerequisites

Please refer to the [Pulumi repo](https://github.com/pulumi/pulumi/)'s [CONTRIBUTING.md file](
https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md#developing) for details on how to get set up with Pulumi for development.

## Committing Generated Code

You must generate and check in the SDKs on each pull request containing a code change, e.g. adding a new resource to `resources.go`.

1. Run `make build_sdks` from the root of this repository
1. Open a pull request containing all changes
1. *Note:* If a large number of seemingly-unrelated diffs are produced by `make build_sdks` (for example, lots of changes to comments unrelated to the change you are making), ensure that the latest dependencies for the provider are installed by running `go mod tidy` in the `provider/` directory of this repository.

## Running Integration Tests

The examples and integration tests in this repository will create and destroy real
cloud resources while running. Before running these tests, make sure that you have
configured access to Checkly with Pulumi.
