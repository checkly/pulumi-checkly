**⚠️ This project is still in very early stages and breaking change could happen**

<p align="center">
  <img width="400px" src="./assets/pulumi.svg" alt="Pulumi" />
</p>

<p>
  <img height="128" src="./assets/checkly.svg" align="right" />
  <h1>Checkly Pulumi Provider</h1>
</p>

> 🟪 Pulumi provider for the [Checkly](https://checklyhq.com) Delightful Active Monitoring

<br>

## 🪛 Installing

This package is only available for JavaScript and TypeScript but support for other languages/platforms, will be available soon.

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
> TBA

### Go
> TBA

### .NET
> TBA

<br>

## 🔑 Authentication

The Pulumi Checkly Provider needs to be configured with Checkly `API Key` and `Account ID` before it can be used to create resources.

> If you don't have and `API Key`, you can create one [here](https://app.checklyhq.com/settings/user/api-keys).

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

<br>


## 🦝 Creating Resources

```javascript
const checkly = require("@checkly/pulumi")

new checkly.Check("api-check", {
  activated: true,
  frequency: 10,
  type: "API",
  request: {
    method: "GET",
    url: "https://checklyhq.com",
  }
})

new checkly.Check("brwoser-check", {
  activated: true,
  frequency: 10,
  type: "BROWSER",
  script: 'console.log("Hello World!")'
})
```

> Check the `examples` directory for more detailed code samples.

## ⚙️  Configuration

The following configuration points are available for the `foo` provider:

- `checkly:apiKey` (environment: `CHECKLY_API_KEY`) - the Checkly API Key.
- `checkly:accountId` (environment: `CHECKLY_ACCOUNT_ID`) - the Checkly account ID.

<br>

## 📖 Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/checkly/api-docs/).

<br>

## 📄 License

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
