const { chromium } = require('playwright')

async function run () {
  const browser = await chromium.launch()
  const page = await browser.newPage()

  const response = await page.goto('https://google.com')

  if (response.status() > 399) {
    throw new Error('Failed with response code ${response.status()}')
  }

  await page.screenshot({ path: 'screenshot.jpg' })

  await page.close()
  await browser.close()
}

run()
