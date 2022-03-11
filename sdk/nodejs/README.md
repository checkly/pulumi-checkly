**⚠️ This project is still in very early stages and breaking changes could happen**

# Checkly Pulumi Provider

The Checkly Pulumi provider enables you to create and configure Checkly resources using your favourite programming language.

## Installation

1. To use this package, please install the Pulumi CLI first.
2. This package is only available for JavaScript and TypeScript but support for other languages/platforms, will be available soon.

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
> TBA

### Go
> TBA

### .NET
> TBA

## Authentication

The Pulumi Checkly Provider needs to be configured with a Checkly `API Key` and `Account ID` before it can be used to create resources.

> If you don't have an `API Key`, you can create one [here](https://app.checklyhq.com/settings/user/api-keys).

Once you generated the `API Key` there are two ways to communicate your authorization tokens to Pulumi:

1. Set the environment variables `CHECKLY_API_KEY` and `CHECKLY_ACCOUNT_ID`:
    ```bash
    $ export CHECKLY_API_KEY=cu_xxx
    $ export CHECKLY_ACCOUNT_ID=xxx
    ```

2. Set them using `pulumi config` command, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:
    ```bash
    $ pulumi config set checkly:apiKey cu_xxx --secret
    $ pulumi config set checkly:accountId xxx
    ```

> Remember to pass `--secret` when setting `checkly:apiKey` so it is properly encrypted.

## Creating Resources

```javascript
const checkly = require("@checkly/pulumi")

new checkly.Check("api-check", {
  activated: true,
  frequency: 10,
  type: "API",
  request: {
    method: "GET",
    url: "https://api.spacexdata.com/v3",
  }
})

new checkly.Check("browser-check", {
  activated: true,
  frequency: 10,
  type: "BROWSER",
  script: `console`
})
```

> Check the `examples` directory for more detailed code samples.

## Configuration

The following configuration points are available for the `foo` provider:

- `checkly:apiKey` (environment: `CHECKLY_API_KEY`) - the Checkly API Key.
- `checkly:accountId` (environment: `CHECKLY_ACCOUNT_ID`) - the Checkly account ID.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/checkly/api-docs/).

## License

[MIT](https://github.com/checkly/pulumi-checkly/blob/main/LICENSE)

<br>


<p align="center">
  <a href="https://checklyhq.com?utm_source=github&utm_medium=sponsor-logo-github&utm_campaign=headless-recorder" target="_blank">
  <img width="100px" src="https://www.checklyhq.com/images/text_racoon_logo.svg" alt="Checkly" />
  </a>
  <br />
  <i><sub>Delightful Active Monitoring for Developers</sub></i>
  <br>
  <b><sub>From Checkly with ♥️</sub></b>
<p>
