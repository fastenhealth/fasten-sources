import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import { generateSourceAuthorizeUrl } from '../../src/connect/authorization-url';
import process from 'process';

test("Flatiron Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);
        let endpointDefinition = await getEndpointDefinition('22d713f7-c8e5-4e5d-98ea-9cdbe9cfe3fb')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        console.log(authorizeData.url.toString());

        await page.goto(authorizeData.url.toString());

        // // Start login process
        await page.waitForSelector("text=Log in");
        await page.click("label[for='Email']", { force: true });
        await page.keyboard.type(process.env.PW_FLATIRON_USERNAME);
        await page.click("label[for='Password']", { force: true });
        await page.keyboard.type(process.env.PW_FLATIRON_PASSWORD);
        await page.click("input[type='submit']");


        await page.waitForSelector("text=This third-party app is requesting access to your information");
        await page.click("button[type='button']");

        await page.waitForSelector("text=Grant Access?");
        // await page.click("div[data-test='atrium-group'] button:nth-of-type(2)");
        await page.click("div[data-test='atrium-group'] button:has-text('Grant Access')");



        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
