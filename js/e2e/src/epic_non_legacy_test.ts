import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import { generateSourceAuthorizeUrl } from '../../src/connect/authorization-url';

test("Epic Non Legacy Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);
        //get the Epic Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('8e2f5de7-46ac-4067-96ba-5e3f60ad52a4')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on MyChart login page
        await page.waitForSelector("text=MyChart Username");
        await expect(page).toHaveTitle("MyChart - Login Page");
        await page.click("label[for='Login']", { force: true });
        await page.keyboard.type("fhircamila");
        await page.click("label[for='Password']", { force: true });
        await page.keyboard.type("epicepic1");
        await page.click("text=Sign In");

        // // We have logged in to MyChart
        await page.waitForSelector("text=Fasten Health has said that it:");
        await page.locator('text=Continue') //wait for continue button
        await expect(page).toHaveTitle("MyChart - Are you sure?");
        await page.getByTitle("Continue to next page").click();


        // We are on the MyChart authorize page. Authorize our app for 1 hour.
        await page.waitForSelector("text=What would you like to share?");
        await expect(page).toHaveTitle("MyChart - Are you sure?");
        await page.click('text=3 months', { force: true, delay: 1000 });
        await page.click("text=Allow access", { force: true, delay: 500 });

        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
