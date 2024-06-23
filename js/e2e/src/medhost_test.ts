import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';
import process from 'process';

test("MedHost Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the Medhost Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('9e5d5b7a-880f-481b-8ae4-77a3d24cfa49')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=YourCare Everywhere");
        await page.focus("#j_username");
        await page.keyboard.type(process.env.PW_MEDHOST_USERNAME);
        await page.focus("#j_password");
        await page.keyboard.type(process.env.PW_MEDHOST_PASSWORD);
        await page.click('#signin');

        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
