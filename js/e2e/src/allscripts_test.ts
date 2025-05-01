import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';

test("Allscripts Login Flow", async ({page}, testInfo) => {
    try{
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);

        //get the Allscripts - Veradigm Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('7682675b-8247-4fda-b2cd-048bfeafc8af')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Allscripts Health Connect Core");
        await expect(page).toHaveTitle("Allscripts FHIR Authorization - ");
        await page.focus("#username");
        await page.keyboard.type("allison.allscripts@tw181unityfhir.edu");
        await page.focus("#passwordEntered");
        await page.keyboard.type("Allscripts#1");
        await page.click("#local-login");

        // We have logged in
        await page.waitForSelector("text=Uncheck the permissions you do not wish to grant.");
        await expect(page).toHaveTitle("Allscripts FHIR Authorization - ");
        await page.click('button[value="yes"]');

        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");

        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
