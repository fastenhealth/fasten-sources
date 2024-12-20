import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import { generateSourceAuthorizeUrl } from '../../src/connect/authorization-url';

test("CareEvolution Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);

        //get the CareEvolution Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('8b47cf7b-330e-4ede-9967-4caa7be623aa')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("#login-button");
        await page.focus("#Username");
        await page.keyboard.type("CEPatient");
        await page.focus("#Password");
        await page.keyboard.type("CEPatient2018");
        await page.click("#login-button");

        // We have logged in
        await page.waitForSelector("text=Whose record do you want to allow access to?");
        // await expect(page).toHaveText("Authorization Request");
        await page.locator("text=Fran Demoski (29)").click();
        await page.click('input[value="Agree"]');

        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
