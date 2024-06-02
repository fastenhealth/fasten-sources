import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';

test.skip("Kaiser Login Flow", async ({page}) => {
    //get the Cerner Sandbox endpoint definition
    let endpointDefinition = await getEndpointDefinition('9d0fa28a-0c5b-4065-9ee6-284ec9577a57')
    let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

    // authorizeData.sourceState
    console.log(authorizeData.url.toString())

    // Start login flow by clicking on button with text "Login to MyChart"
    await page.goto(authorizeData.url.toString());

    // We are on login page
    await page.waitForSelector("text=Sign in");
    await expect(page).toHaveTitle("Sign in ");
    await page.click("label[for='username']", { force: true });
    await page.keyboard.type("Pvaluser1");
    await page.click("label[for='password']", { force: true });
    await page.keyboard.type("V@lidation1");
    await page.click("button[type='submit']");

    // We have logged in
    await page.waitForSelector("text=Please Be Aware");
    await expect(page).toHaveTitle("Consent Management");
    await page.click("input[name='subject']");
    await page.click("label[for='agreement']");
    await page.locator('div.accept-btn').locator('button[type="button"]').click();


    // If successful, Fasten Lighthouse page should now be visible
    await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
});
