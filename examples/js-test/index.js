
const  checkly = require( "@pulumi/checkly");
const  pulumi = require( "@pulumi/pulumi");

const x = new checkly.Check("pulumi-test", {
  name: "testCheck",
  activated: true,
  frequency: 10,
  type: "API",
  request: {
    method: "GET",
    url: "https://checklhq.com",
  }
});

new checkly.CheckGroup("pulumi-group", {
  name: "Pulumi Group",
  activated: true,
  concurrency: 1,
  locations: ['us-east-1'],
  apiCheckDefaults: {
    url: 'https://www.pulumi.com',
  },
});

