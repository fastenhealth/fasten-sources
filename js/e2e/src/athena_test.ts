import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';

test("Athena Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the Cerner Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('950e9092-8ce7-4926-ad87-64616f00cb4c')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Fasten Health - preview");
        await expect(page).toHaveTitle("Login");
        await page.click("label[for='okta-signin-username']", { force: true });
        await page.keyboard.type("phrtest_preview@mailinator.com");
        await page.click("label[for='okta-signin-password']", { force: true });
        await page.keyboard.type("Password1");
        await page.click("#okta-signin-submit");

        // We have logged in
        await page.waitForSelector("text=Select a health record");
        await expect(page).toHaveTitle("Login");
        await page.locator("text=Jake Medlock (you)").click();

        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");

        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
