<p>
  <img height="128" src="https://www.checklyhq.com/images/footer-logo.svg" align="right" />
  <h1>Checkly Pulumi Provider</h1>
</p>

![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/checkly/pulumi-checkly?label=Version)

The Checkly Pulumi provider enables you to create and configure Checkly resources using your favourite programming language.
Note that this project is in its early stages and breaking changes could happen.

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/get-started/install/).
2. This package is only available for JavaScript and TypeScript but support for other languages will be available soon.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm` or `yarn`:

```bash
$ npm install @checkly/pulumi
$ yarn add @checkly/pulumi
```

### Python, Go & .NET

*TBA*

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

## Getting Started

1. Open your terminal and run `$ pulumi new javascript` to create a new Pulumi project with the `javascript` template.
1. Install the Checkly Pulumi provider using npm: `$ npm i @checkly/pulumi`.
1. Look for `index.js` file in the root of your project and replace content with the following code:

    ```javascript
    const checkly = require("@checkly/pulumi")

    new checkly.Check("api-check", {
      type: "API",
      name: "Public SpaceX API",
      activated: true,
      frequency: 10,
      locations: ["us-east-1"],
      request: {
        method: "GET",
        url: "https://api.spacexdata.com/v3",
        assertions: [
          {
            source: 'STATUS_CODE',
            comparison: 'EQUALS',
            target: 200
          },
          {
            source: 'JSON_BODY',
            property: '$.project_name',
            comparison: 'EQUALS',
            target: 'SpaceX-API'
          }
        ]
      }
    })

    new checkly.Check("browser-check", {
      type: "BROWSER",
      name: "Google.com Playwright check",
      activated: true,
      frequency: 10,
      locations: ["us-east-1"],
      script: `const { chromium } = require('playwright')

    async function run () {
      const browser = await chromium.launch()
      const page = await browser.newPage()

      const response = await page.goto('https://google.com')

      if (response.status() > 399) {
        throw new Error('Failed with response code \${response.status()}')
      }

      await page.screenshot({ path: 'screenshot.jpg' })

      await page.close()
      await browser.close()
    }

    run()`
    })
    ```
1. Setup you Checkly API Key and Account id:
    ```bash
    $ pulumi config set checkly:apiKey cu_xxx --secret
    $ pulumi config set checkly:accountId xxx
    ```
1. You are ready to go, run `$ pulumi up` to deploy your stack 🚀

## Examples

Find working JavaScript and TypeScript code samples in the [`./examples`](https://github.com/checkly/pulumi-checkly/tree/main/examples) directory.

## Learn More
For documentation and example usage see:
1. [Checkly's documentation](https://www.checklyhq.com/docs/integrations/pulumi/).
2. [The official provider documentation](https://www.pulumi.com/registry/packages/checkly/api-docs/)
3. [Working Examples](https://github.com/checkly/pulumi-checkly/tree/main/examples).

## Questions
For questions and support please open a new  [discussion](https://github.com/checkly/pulumi-checkly/discussions). The issue list of this repo is exclusively for bug reports and feature/docs requests.

## Issues
Please make sure to respect issue requirements and choose the proper [issue template](https://github.com/checkly/pulumi-checkly/issues/new/choose) when opening an issue. Issues not conforming to the guidelines may be closed.

## Contribution
Please make sure to read the [Contributing Guide](https://github.com/checkly/pulumi-checkly/blob/main/CONTRIBUTING.md) before making a pull request.

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
