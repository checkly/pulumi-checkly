# Checkly Provider for Pulumi

![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/checkly/pulumi-checkly?label=Version)

The Checkly provider for Pulumi can be used to provision any of the monitoring resources available in [Checkly](https://www.checklyhq.com/).

## Installation

The Checkly provider is available as a package in most Pulumi languages:

* JavaScript/TypeScript: [`@checkly/pulumi`](https://www.npmjs.com/package/@checkly/pulumi)
* Python: [`pulumi-checkly`](https://pypi.org/project/pulumi-checkly/)
* Go: [`github.com/checkly/pulumi-checkly/sdk/v2/go/checkly`](https://github.com/checkly/pulumi-checkly)
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

## Configuration options

The following configuration points are available for the Checkly provider:

- `checkly:accountId` (environment: `CHECKLY_ACCOUNT_ID`) - your Checkly Account ID
- `checkly:apiKey` (environment: `CHECKLY_API_KEY`) - your Checkly API Key
    * If you don't have an API Key, you can create one [here](https://app.checklyhq.com/settings/user/api-keys).
    * Make sure to use the `--secret` flag with `pulumi config set`.
- `checkly:apiUrl` (environment: `CHECKLY_API_URL`) - for internal development purposes only

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/checkly/api-docs/).
