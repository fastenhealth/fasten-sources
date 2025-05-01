import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import { generateSourceAuthorizeUrl } from '../../src/connect/authorization-url';

test("Cerner Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);

        //get the Cerner Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('3290e5d7-978e-42ad-b661-1cf8a01a989c')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=FhirPlay Non-Prod");
        await expect(page).toHaveTitle("Cerner Health - Sign In");
        await page.click("label[for='id_login_username']", { force: true });
        await page.keyboard.type("nancysmart");
        await page.click("label[for='id_login_password']", { force: true });
        await page.keyboard.type("Cerner01");
        await page.click("#signin");

        // We have logged in
        await page.waitForSelector("text=Warning: Unknown app");
        await expect(page).toHaveTitle("Authorization Needed");
        await page.click('#proceedButton');

        // We are on the Select Patient page.
        await page.waitForSelector("text=Allow access to the record(s) of:");
        await expect(page).toHaveTitle("Authorization Needed");
        // await page.click("label[for='12724066']", { force: true, delay: 500 });
        await page.click("#allowButton");


        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");

        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
