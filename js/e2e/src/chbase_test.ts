import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';
import process from 'process';

test("CHbase Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);
        //get the VAHealth Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('ee5e19b6-4539-4e46-baab-b892061fe448')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // // We are on landing page
        await page.waitForSelector("text=Sign in to your account");
        await expect(page).toHaveTitle("Sign in to Unify Login");

        // // We are on the ID.me login page
        // await page.waitForSelector("text=Sign in to ID.me");
        await page.click("label[for='username']", { force: true });
        await page.keyboard.type(process.env.PW_CHBASE_USERNAME);
        await page.click("label[for='password']", { force: true });
        await page.keyboard.type(process.env.PW_CHBASE_PASSWORD);
        await page.click("input[type='submit']");

        await page.waitForSelector('#RecordFacilityList');

        // Select the option by its value
        await page.selectOption('#RecordFacilityList', 'ae1dc61d-e73d-4d85-a273-2f1b0f22b24c');
        await page.click("button[type='submit']");


        await page.waitForSelector("text=Choose what data to share");
        await page.click("input[type='submit']");

        // // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
