const puppeteer = require('puppeteer')
const browser = await puppeteer.launch()
const page = await browser.newPage()
const navigationPromise = page.waitForNavigation()

await page.goto('https://www.checklyhq.com/')

await page.setViewport({ width: 1792, height: 1016 })

await page.waitForSelector('.left-side > .dropdown-item:nth-child(1) > .d-flex > div > .menu-title')
await page.click('.left-side > .dropdown-item:nth-child(1) > .d-flex > div > .menu-title')

await navigationPromise

await browser.close()
