import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';

test.skip("Cigna Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the Cerner Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('6c0454af-1631-4c4d-905d-5710439df983')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Log in to share health information");
        await expect(page).toHaveTitle("myCigna - Auth Flow");
        await page.click("label[for='username']", { force: true });
        await page.keyboard.type("syntheticuser05");
        await page.click("label[for='password']", { force: true });
        await page.keyboard.type("5ynthU5er5");
        await page.click("button[type='submit']");

        // We have logged in
        await page.waitForSelector("label[for='termsAccept']");
        await expect(page).toHaveTitle("myCigna - Auth Flow");
        await page.click('label[for="termsAccept"]', { force: true });
        await page.click('button[type="submit"]');

        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
