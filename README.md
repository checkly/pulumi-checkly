**🚨 This project is still in very early stages and is not stable, _use at your own risk_! 🚨**

# Checkly Pulumi Provider

The Checkly Resource Provider lets you manage [Checkly](https://checklyhq.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/checkly
```

or `yarn`:

```bash
yarn add @pulumi/checkly
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_checkly
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-checkly/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Checkly
```

## Configuration

The following configuration points are available for the `foo` provider:

- `checkly:apiKey` (environment: `CHECKLY_API_KEY`) - the Checkly API Key.
- `checkly:accountId` (environment: `CHECKLY_ACCOUNT_ID`) - the Checkly account ID.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/).
