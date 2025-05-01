import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';
import process from 'process';

test("DrChrono Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the DrChrono Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('6a01ffff-d73e-4728-a315-9b23bd77a4cc')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Sign In");
        await page.focus("#username");
        await page.keyboard.type(process.env.PW_DRCHRONO_USERNAME);
        await page.focus("#password");
        await page.keyboard.type(process.env.PW_DRCHRONO_PASSWORD);
        await page.click('button:text("Sign In")');

        // We have logged in
        await page.waitForSelector("text=Charlie DeBraga");
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
