import * as pulumi from "@pulumi/pulumi";
import * as checkly from "@pulumi/checkly";

import * as fs from "fs";

new checkly.Check("pulumi-test", {
    name: "Pulumi Test",
    activated: true,
    frequency: 10,
    type: "API",
    request: {
      method: "GET",
      url: "https://checklhq.com",
    }
});


const script = fs.readFileSync('./test.js', 'utf8')

new checkly.Check("pulumi-browser", {
  name: "Pulumi Browser",
  activated: true,
  frequency: 10,
  type: "BROWSER",
  script,
});


