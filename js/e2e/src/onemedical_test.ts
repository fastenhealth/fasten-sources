import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';
import process from 'process';

test("OneMedical Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the OneMedical Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('c835742c-c896-4b93-beb5-28df18f16bd8')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Log in to continue");
        await page.focus("#email");
        await page.keyboard.type(process.env.PW_ONEMEDICAL_USERNAME);
        await page.focus("#password");
        await page.keyboard.type(process.env.PW_ONEMEDICAL_PASSWORD);
        await page.click('#btn-login');

        // We have logged in
        await page.waitForSelector("text=Connect Your Records");
        await page.click('#scope-selection-submit');

        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
