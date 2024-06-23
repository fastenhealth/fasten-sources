import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';
import process from 'process';

test("MaximEyes Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the Maximeyes Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('f4ea79fc-9a1b-43f0-be95-24247d666514')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Forgot my password");
        await page.focus("#Username");
        await page.keyboard.type(process.env.PW_MAXIMEYES_USERNAME);
        await page.focus("#Password");
        await page.keyboard.type(process.env.PW_MAXIMEYES_PASSWORD);
        await page.click('input[value="Sign in"]');

        // We have logged in
        await page.waitForSelector("text=Reed");
        await page.click("input[value='Select']");

        // We have logged in
        await page.waitForSelector("text=Access");
        await page.click("input[value='Yes']");


        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
