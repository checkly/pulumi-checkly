# Checkly Provider for Pulumi

![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/checkly/pulumi-checkly?label=Version)

The Checkly provider for Pulumi can be used to provision any of the monitoring resources available in [Checkly](https://www.checklyhq.com/).

## Installation

The Checkly provider is available as a package in most Pulumi languages:

* JavaScript/TypeScript: [`@checkly/pulumi`](https://www.npmjs.com/package/@checkly/pulumi)
* Python: [`pulumi-checkly`](https://pypi.org/project/pulumi-checkly/)
* Go: [`github.com/checkly/pulumi-checkly/sdk/v3/go/checkly`](https://github.com/checkly/pulumi-checkly)
* .NET: [`Pulumi.Checkly`](https://www.nuget.org/packages/Pulumi.Checkly)


## Authentication

The Checkly provider must be configured with an `API Key` and an `Account ID` in order to deploy Checkly resources. Sign up for a [Checkly](https://www.checklyhq.com) account and follow our [integration guide](https://www.checklyhq.com/docs/integrations/pulumi/) to create and configure your credentials.

### Example configuration

First, configure your Checkly Account ID:

```
pulumi config set checkly:accountId YOUR_CHECKLY_ACCOUNT_ID
```

Then, configure you Checkly API key (with `--secret`):

```
pulumi config set checkly:apiKey YOUR_CHECKLY_API_KEY --secret
```

You should now be able to deploy Checkly resources.

## Example usage

You can find working JavaScript and TypeScript code samples in the [`./examples`](https://github.com/checkly/pulumi-checkly/tree/main/examples) directory.

## Upgrading to v3

v3 bridges version 2 of the Checkly Terraform provider, which moved to the Terraform Plugin Framework. This section lists what changes for Pulumi programs and stacks. For the details behind it, and for smaller changes not listed here, read the [Terraform provider's v2 upgrade guide](https://registry.terraform.io/providers/checkly/checkly/latest/docs/guides/upgrading-to-v2).

### Upgrade procedure

1. While still on v2.x, run `pulumi up` and apply any pending changes, so that you start from a clean preview.
1. Delete every `TriggerCheck` and `TriggerCheckGroup` resource while still on v2.x: remove it from your program and run `pulumi up`. To keep the trigger URL working for now, run `pulumi state delete <urn>` instead, which removes the resource from the stack only. If you have already upgraded, use `pulumi state delete <urn>`: v3 doesn't have these resource types and can't delete them.
1. Back up the stack with `pulumi stack export --file backup.json`. Once v3 has updated or refreshed the stack, v2.x may not be able to read it; to go back, restore the backup with `pulumi stack import --file backup.json`.
1. Make sure every API key has an account ID: set `checkly:accountId` or the `CHECKLY_ACCOUNT_ID` environment variable. v2.x only required it for user-scoped keys.
1. Upgrade the package to v3 and update your program as described in [Changes to your program](#changes-to-your-program).
1. Run `pulumi preview`, then `pulumi up`. The preview should show only in-place updates, where a default changed; see [Expected changes](#expected-changes). If it shows a deletion or replacement you didn't intend, fix your program before running `pulumi up`. This first `pulumi up` switches the stack to the v3 provider. Don't run `pulumi refresh --run-program` before it: it fails on the inputs that v2.x stored, with `objectEncoder failed on property "alert_settings"`. A plain `pulumi refresh` before it still uses the v2.x provider recorded in the stack.
1. Optionally, move deprecated resources to their replacements, as described in [Moving to the replacement resources](#moving-to-the-replacement-resources).

### Changes to your program

- Go: the import path is now `github.com/checkly/pulumi-checkly/sdk/v3/go/checkly`.
- Go: the SDK needs Go 1.26.6 or later and `github.com/pulumi/pulumi/sdk/v3` v3.276.0 or later.
- `TriggerCheck` and `TriggerCheckGroup` are gone. Checkly has deprecated command line trigger URLs; run checks on demand with the [Checkly CLI](https://www.checklyhq.com/docs/cli/) instead.
- New resources: `ApiCheck`, `BrowserCheck` and `MultistepCheck`, one per check type, which expose only the properties of their type. They replace `Check`.
- `MaintenanceWindow` now manages the upstream `checkly_maintenance_window` resource. Your code keeps working unchanged.
- Deprecated, but still available: `Check`, `CheckGroup` (use `CheckGroupV2`), `HeartbeatCheck` (use `HeartbeatMonitor`), `TcpCheck` (use `TcpMonitor`), `StatusPage` (use `StatusPageV3`) and `StatusPageService` (use `StatusPageV3Component`).
- Nested `alertSettings` values are objects instead of lists, with singular names: `runBasedEscalation`, `timeBasedEscalation`, `parallelRunFailureThreshold` and `reminders`. `alertSettings.sslCertificates` is gone. This applies to every check, group and monitor, and to `CheckGroupV2.enforceAlertSettings.alertSettings`.
- `Check` and `CheckGroup`: `environmentVariables` is a list of `{ key, value, locked, secret }` objects. It replaces both the `environmentVariable` list and the `environmentVariables` map; rewrite each map entry as a list element with the same `key` and `value`. Remove environment variables from API checks, where they never had an effect.
- `groupId`, `setupSnippetId` and `teardownSnippetId` must be left out, or set to `undefined` (`None` in Python, `nil` in Go), when a check has no group or snippet. `0` now fails the preview. An expression that chooses between an ID and none must choose `undefined` instead of `0`.
- `Check`: `sslCheck` is gone. On a browser or multi-step `Check`, setting `shouldFail` fails the preview, even to `false`; remove it.
- `GrpcMonitor`, `SslMonitor` and `TracerouteMonitor`: `shouldFail` is gone.
- `DnsMonitor`, `GrpcMonitor`, `IcmpMonitor`, `SslMonitor`, `TcpCheck`, `TcpMonitor` and `TracerouteMonitor`: `retryStrategy.onlyOn` is gone.
- `CheckGroup` and `CheckGroupV2`: when you set `apiCheckDefaults`, its `url` is required. Leave `apiCheckDefaults` out if you don't need a base URL.
- These outputs are now secrets: environment variable values (`EnvironmentVariable.value` and `environmentVariables[].value` on checks, check groups and Playwright check suites), basic auth passwords, `heartbeat.pingToken`, `PrivateLocation.proxyUrl`, and `slack.url`, `webhook.url`, `webhook.webhookSecret`, `opsgenie.apiKey` and `pagerduty.serviceKey` of `AlertChannel`. Values derived from them, such as a ping URL built from `pingToken`, become secrets too; `pulumi stack output --show-secrets` shows them.
- The provider checks lengths and formats in the preview, so some values that the API silently changed or dropped now fail. The upstream guide lists them under [Values the API rejects or changes](https://registry.terraform.io/providers/checkly/checkly/latest/docs/guides/upgrading-to-v2#values-the-api-rejects-or-changes).
- `CHECKLY_DEBUG_LOG` no longer writes API traffic to a file. Use an HTTPS proxy through `HTTPS_PROXY` instead.

### Expected changes

The first `pulumi up` on v3 can update resources in place. The most important cases are:

- `useGlobalAlertSettings` no longer has a fixed default: it is `true` when `alertSettings` isn't set, and `false` when it is. Checks, check groups, heartbeats and the TCP, gRPC, SSL and traceroute monitors that set neither switch to the account level alert settings; set `useGlobalAlertSettings: false` to keep their own. `UrlMonitor`, `DnsMonitor`, `IcmpMonitor` and `PlaywrightCheckSuite` that set `alertSettings` but not `useGlobalAlertSettings` start using their `alertSettings`; set `useGlobalAlertSettings: true` and remove `alertSettings` to keep the account level settings. On `CheckGroupV2`, the same applies to `enforceAlertSettings` while it is enabled.
- Default response time limits match the API: `SslMonitor`, `TcpMonitor`, `TcpCheck` and `TracerouteMonitor` that leave them out get lower limits. Set them to keep the old ones.
- `AlertChannel.autoSubscribe` is new and defaults to `false`, which turns off automatic subscription of new checks. Set it to `true` to keep it on.
- Some properties that v2.x couldn't clear, such as a check's `tags` or a dashboard's `logo`, are now cleared when you leave them out. Add values set in the web app to your program to keep them.

The upstream guide lists every case under [Expected one-time changes](https://registry.terraform.io/providers/checkly/checkly/latest/docs/guides/upgrading-to-v2#expected-one-time-changes).

### Moving to the replacement resources

You can move a deprecated resource to its replacement without recreating it. After your first `pulumi up` on v3, run `pulumi refresh` once, so that the stored state of the old resource is in the v3 shape. Then change the resource in your program, add an alias with the old type, and run `pulumi up`:

```typescript
new checkly.ApiCheck("my-api-check", {
  // The properties of the check, with the type property removed.
}, { aliases: [{ type: "checkly:index/check:Check" }] })
```

The resource is updated in place. This works for:

- `Check` to `ApiCheck`, `BrowserCheck` or `MultistepCheck`. Pick the resource that matches the check's `type`: the provider can't verify it during the move.
  - `ApiCheck` defaults `degradedResponseTime` and `maxResponseTime` to 5000 and 20000 instead of 15000 and 30000, and `request.followRedirects` to `true` instead of `false`. Set them to keep the previous behaviour.
  - The per-type resources have no `doubleCheck`. Before moving a check that sets it, replace it with the equivalent `retryStrategy` (type `FIXED`, `maxRetries: 1`, `baseBackoffSeconds: 0`, `maxDurationSeconds: 600`, `sameRegion: false`), or the check stops retrying.
  - Before moving a browser or multi-step check whose `shouldFail` is `true` in Checkly, remove `shouldFail` from the `Check` and run `pulumi up`, which sets it to `false`.
- `CheckGroup` to `CheckGroupV2`. A v2 group doesn't impose settings on its checks unless you enforce them, so read [Check groups](https://registry.terraform.io/providers/checkly/checkly/latest/docs/guides/upgrading-to-v2#check-groups) in the upstream guide first.
- `HeartbeatCheck` to `HeartbeatMonitor`, and `TcpCheck` to `TcpMonitor`.

`StatusPage` and `StatusPageService` can't be moved: v3 status pages are separate objects in Checkly. Keep them, or create a `StatusPageV3` page with `StatusPageV3Component` resources and then delete the old ones.

## Configuration options

The following configuration points are available for the Checkly provider:

- `checkly:accountId` (environment: `CHECKLY_ACCOUNT_ID`) - your Checkly Account ID
- `checkly:apiKey` (environment: `CHECKLY_API_KEY`) - your Checkly API Key
    * If you don't have an API Key, you can create one [here](https://app.checklyhq.com/settings/user/api-keys).
    * Make sure to use the `--secret` flag with `pulumi config set`.
- `checkly:apiUrl` (environment: `CHECKLY_API_URL`) - for internal development purposes only

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/checkly/api-docs/).
