import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';
import process from 'process';

test("Practice Fusion Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the Practice Fusion Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('8b340f03-7643-4315-9bb6-6b01e31e5402')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on scope page
        await page.waitForSelector("text=This application would like access to your information");
        await page.click('button:text("Allow")');

        // We are on login page
        await page.waitForSelector("text=Don't remember your password?");
        await page.focus("input[name='username']");
        await page.keyboard.type(process.env.PW_PRACTICEFUSION_USERNAME);
        await page.focus("input[name='password']");
        await page.keyboard.type(process.env.PW_PRACTICEFUSION_PASSWORD);
        await page.click("button[name='submit']");


        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
