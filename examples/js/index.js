const fs = require('fs')
const checkly = require( "@pulumi/checkly");

const script = fs.readFileSync("./test.js", 'utf-8');

const channel = new checkly.AlertChannel("pulumi-channel", {
  email: {
    address: 'nacho@checklyhq.com'
  }
})

const group = new checkly.CheckGroup("pulumi-group", {
  name: "Pulumi Group",
  activated: true,
  concurrency: 1,
  locations: ['us-east-1'],
  apiCheckDefaults: { url: 'https://www.pulumi.com' },
  alertChannelSubscriptions: [
    {
      activated: true,
      channelId: channel.id.apply(id => parseInt(id)),
    }
  ]
});

new checkly.Check("pulumi-api-check", {
  activated: true,
  frequency: 10,
  type: "API",
  groupId: group.id.apply(id => parseInt(id)),
  request: {
    method: "GET",
    url: "https://checklhq.com",
  }
});

new checkly.Check("pulumi-brwoser-check", {
  activated: true,
  frequency: 10,
  type: "BROWSER",
  groupId: group.id.apply(id => parseInt(id)),
  script,
});

new checkly.Snippet('snippet', {
  script,
})