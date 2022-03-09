const fs = require('fs')
const checkly = require("@pulumi/checkly")

const snippetScript = fs.readFileSync("./scripts/snippet.js", 'utf-8')
const browerCheckScript = fs.readFileSync("./scripts/browser-check.js", 'utf-8')

const PREFIX = 'pulimi-js'


const emailChannel = new checkly.AlertChannel(PREFIX + "email-channel", {
  email: {
    address: 'your@email.com'
  }
})

const slackChannel = new checkly.AlertChannel(PREFIX + "slack-channel", {
  slack: {
    url: 'https://hooks.slack.com/services/X0JKQJQ0P',
    channel: '#alerts',
  }
})

const group = new checkly.CheckGroup(PREFIX + "group", {
  activated: true,
  concurrency: 1,
  locations: ['us-east-1', 'us-west-1'],
  apiCheckDefaults: {},
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

new checkly.Check(PREFIX + "api-check", {
  activated: true,
  frequency: 10,
  type: "API",
  groupId: group.id.apply(id => parseInt(id)),
  request: {
    method: "GET",
    url: "https://checklyhq.com",
  }
})

new checkly.Check(PREFIX + "brwoser-check", {
  activated: true,
  frequency: 10,
  type: "BROWSER",
  groupId: group.id.apply(id => parseInt(id)),
  script: browerCheckScript
})

new checkly.Snippet('snippet', { script: snippetScript})

new checkly.MaintenanceWindow(PREFIX + 'maintenance', {
  startsAt: '2022-03-01',
  endsAt: '2022-03-02',
  repeatInterval: 1,
  repeatUnit: 'DAY',
})

new checkly.TriggerCheckGroup(PREFIX + 'trigger', {
groupId: group.id.apply(id => parseInt(id)),
})

new checkly.PublicDashboard(PREFIX + 'dashboard', {
  customUrl: PREFIX + 'js',
  paginationRate: 30,
  refreshRate: 300,
})