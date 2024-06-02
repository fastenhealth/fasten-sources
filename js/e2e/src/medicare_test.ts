import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';

test.skip("Medicare Login Flow", async ({page}) => {
    //get the Cerner Sandbox endpoint definition
    let endpointDefinition = await getEndpointDefinition('6ae6c14e-b927-4ce0-862f-91123cb8d774')
    let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

    // authorizeData.sourceState
    console.log(authorizeData.url.toString())

    // Start login flow by clicking on button with text "Login to MyChart"
    await page.goto(authorizeData.url.toString());

    // We are on login page
    await page.waitForSelector("text=Log in");
    // await expect(page).toHaveTitle("Cerner Health - Sign In");
    await page.click("label[for='username-textbox']", { force: true });
    await page.keyboard.type("BBUser00000");
    await page.click("label[for='password-textbox']", { force: true });
    await page.keyboard.type("PW00000!");
    await page.click("#login-button");

    //INCOMPLETE - Need to add more steps to complete the login flow

    // // We have logged in
    // await page.waitForSelector("text=Warning: Unknown app");
    // await expect(page).toHaveTitle("Authorization Needed");
    // await page.click('#proceedButton');
    //
    // // We are on the Select Patient page.
    // await page.waitForSelector("text=SMART II, NANCY (Self, 33)");
    // await expect(page).toHaveTitle("Authorization Needed");
    // await page.click("label[for='12724066']", { force: true, delay: 500 });
    // await page.click("#allowButton");
    //
    //
    // // If successful, Fasten Lighthouse page should now be visible
    // await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
});
