const fs = require('fs')
const checkly = require("@checkly/pulumi")

// We store code snippets separately from the Checkly entities.
const snippetScript = fs.readFileSync("./scripts/snippet-script.js", "utf-8")
const playwrightScript = fs.readFileSync("./scripts/playwright-script.js", "utf-8")

// Some variables we can reuse in checks and check groups
const locations = ['eu-west-1', 'us-west-2']
const tags = ['pulumi']

// Alert channels are defined up front so we can add them to checks and check groups later.
const emailChannel = new checkly.AlertChannel("my-email-channel", {
  email: {
    address: 'your@email.com'
  }
})

const slackChannel = new checkly.AlertChannel("my-slack-channel", {
  slack: {
    url: 'https://hooks.slack.com/services/<REPLACE_WITH_ACTUAL_SLACK_HOOK>',
    channel: '#alerts',
  }
})

// We create a group, add the tags, locations and alert channels.
const group = new checkly.CheckGroup("my-group", {
  name: 'Check Group #1',
  activated: true,
  concurrency: 1,
  apiCheckDefaults: {},
  locations,
  tags,
  useGlobalAlertSettings: true,
  alertChannelSubscriptions: [
    {
      activated: true,
      channelId: slackChannel.id.apply(id => parseInt(id)),
    },
    {
      activated: true,
      channelId: emailChannel.id.apply(id => parseInt(id)),
    }
  ]
})

// Our first Browser check is added to group created above and uses the Playwright script.
new checkly.Check("my-browser-check", {
  name: "Google.com Playwright check",
  activated: true,
  frequency: 10,
  type: "BROWSER",
  groupId: group.id.apply(id => parseInt(id)),
  script: playwrightScript,
  locations,
  tags
})

const setupSnippet = new checkly.Snippet('my-snippet', {
  name: 'set-header',
  script: snippetScript
})

// Our first API check uses the setup script defined above and has two assertions.
new checkly.Check("my-api-check", {
  name: "Public SpaceX API",
  activated: true,
  frequency: 10,
  type: "API",
  locations,
  tags,
  setupSnippetId: setupSnippet.id.apply(id => parseInt(id)),
  degradedResponseTime: 5000,
  maxResponseTime: 15000,
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
  },
  useGlobalAlertSettings: true,
  alertChannelSubscriptions: [
    {
      activated: true,
      channelId: emailChannel.id.apply(id => parseInt(id)),
    }
  ]
})

// This maintenance window runs for 1 hour after midnight UTC and repeats daily for two months.
// All the checks with the "tags" array are included in this window.
new checkly.MaintenanceWindow('my-maintenance-window', {
  name: "Daily maintenance",
  startsAt: '2022-03-01T00:00:00.000Z',
  endsAt: '2022-03-01T01:00:00.000Z',
  repeatEndsAt: '2022-05-02T00:00:00.000Z',
  repeatInterval: 1,
  repeatUnit: 'DAY',
  tags
})

// The public dashboard shows all the checks tags with the "tags" array
new checkly.Dashboard('my-dashboard', {
  customUrl: "my-pulumi-checkly-dash",
  customDomain: "",
  header: "Checks created with Pulumi",
  hideTags: false,
  logo: "",
  paginate: false,
  paginationRate: 30,
  refreshRate: 300,
  tags
})

// Environment variables can be used in API request URLs, headers, query parameters etc.
// Use variables in API checks by using the {{MY_VAR}} notation.
// Use variables in Browser checks with the process.env.MY_VAR notation
new checkly.EnvironmentVariable('my-env-var', {
  key: 'API_URL',
  value: 'https://localhost:3000',
})