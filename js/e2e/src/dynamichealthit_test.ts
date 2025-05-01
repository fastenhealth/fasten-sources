import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';
import process from 'process';

test("DynamicHealthIT Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the NextGen Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('5f928520-2b73-4a47-b1a1-30ed6c3d9e9b')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Account Login");
        await page.focus("#username");
        await page.keyboard.type(process.env.PW_DYNAMICHEALTHIT_USERNAME);
        await page.focus("#password");
        await page.keyboard.type(process.env.PW_DYNAMICHEALTHIT_PASSWORD);
        await page.click('button:text("Sign In")');

        // We have logged in
        await page.waitForSelector("text=RebeccaF Larson");
        await page.click('button:text("Select")');

        // Keep me signed in.
        await page.waitForSelector("text=Remember My Decision");
        await page.click('button:text("Yes, Allow")');


        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
