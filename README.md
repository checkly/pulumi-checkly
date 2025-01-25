# Checkly Pulumi Provider

![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/checkly/pulumi-checkly?label=Version)

The Checkly Pulumi Provider lets you manage [Checkly](https://www.checklyhq.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @checkly/pulumi
```

or `yarn`:

```bash
yarn add @checkly/pulumi
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_checkly
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/checkly/pulumi-checkly/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Checkly
```

## Getting Started

Please see our [integration guide](https://www.checklyhq.com/docs/integrations/pulumi/) for a complete setup flow.

## Configuration

The following configuration points are available for the `checkly` provider:

- `checkly:accountId` (environment: `CHECKLY_ACCOUNT_ID`) - your Checkly Account ID
- `checkly:apiKey` (environment: `CHECKLY_API_KEY`) - your Checkly API Key
    * If you don't have an API Key, you can create one [here](https://app.checklyhq.com/settings/user/api-keys).
    * Make sure to use the `--secret` flag with `pulumi config set`.
- `checkly:apiUrl` (environment: `CHECKLY_API_URL`) - for internal development purposes only

## Examples

Find working JavaScript and TypeScript code samples in the [`./examples`](./examples) directory.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/checkly/api-docs/).
