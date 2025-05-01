import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';

test("Nextgen Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the NextGen Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('843f5c82-b4e3-43c6-8657-eff1390d7e44')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=NextGen HealthCare");
        await page.focus("#Username");
        await page.keyboard.type("patientapitest");
        await page.focus("#Password");
        await page.keyboard.type("Password1!");
        await page.click('button:text("Next")');

        // We have logged in
        await page.waitForSelector("text=Connect to Fasten Health - Sandbox");
        await page.click('#btnAllow');

        // Keep me signed in.
        await page.waitForSelector("text=Connect to Fasten Health - Sandbox");
        await page.click('button:text("Next")');


        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
