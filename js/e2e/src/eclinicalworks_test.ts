import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateFastenConnectAuthorizeUrl, generateSourceAuthorizeUrl} from '@shared-library';

test("eClinicalWorks-Healow Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the eClinicalWorks Sandbox endpoint definition
        // let endpointDefinition = await getEndpointDefinition('f0a8629a-076c-4f78-b41a-7fc6ae81fa4d')
        // let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        let authorizeData = await generateFastenConnectAuthorizeUrl(
            '41b42645-350f-475e-b21b-e7a1276fca4f',
            'a1a5bbc0-2e6a-4d6b-bf1c-d4faa4820c77',
            'f0a8629a-076c-4f78-b41a-7fc6ae81fa4d',
        )

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=FHIR R4 Prodtest EMR");
        await expect(page).toHaveTitle("healow - Health and Online Wellness");
        await page.focus("#username");
        await page.keyboard.type("AdultFemaleFHIR");
        await page.focus("#pwd");
        await page.keyboard.type("e@CWFHIR1");
        await page.click("#btnLoginSubmit");

        // We have logged in
        await page.waitForSelector("text=What you need to know about Fasten Health");
        await expect(page).toHaveTitle("LoginUi");
        await page.click('button:text(" Continue ")');


        // We are on the agreement page
        await page.waitForSelector("text=Personal Information Sharing");
        await expect(page).toHaveTitle("LoginUi");
        //this is a case-sensitive, but not exact match. It may break in the future.
        await page.click('button:text-is("Approve")');

        // If successful, redirect page should now be visible
        await page.waitForSelector("text=Example Domain");

        //parse the query string parameters for the current url.
        const url = new URL(page.url());
        const params = new URLSearchParams(url.search);
        //check if the required parameters are present
        // https://www.example.com/?brand_id=5b7ff2c3-804f-4443-9bd9-4437d43c3b87&connection_status=authorized&endpoint_id=3290e5d7-978e-42ad-b661-1cf8a01a989c&external_state=489bfc62-085b-4a9e-8664-6961106d9120&org_connection_id=25aebcd5-5f2b-435d-bbd2-f0878dd1b4b2&platform_type=cerner&portal_id=00a83214-7b14-4a12-ad95-5198b70dbb63&request_id=f37557d8-1696-4ecb-88f7-8138d39282b5
        expect(params.has('org_connection_id')).toBe(true);
        expect(params.has('request_id')).toBe(true);
        expect(params.get('connection_status')).toEqual("authorized");
        expect(params.get('external_state')).toEqual(authorizeData.sourceState.state);

        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
