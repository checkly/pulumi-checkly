
const  checkly = require( "@pulumi/checkly");
const  pulumi = require( "@pulumi/pulumi");

const x = new checkly.Check("pulumi-test", {
  name: "testCheck",
  activated: true,
  frequency: 10,
  type: "API",
});

