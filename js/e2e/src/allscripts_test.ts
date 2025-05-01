import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateFastenConnectAuthorizeUrl, generateSourceAuthorizeUrl} from '@shared-library';

test("Allscripts Login Flow", async ({page}, testInfo) => {
    try{
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);

        //get the Allscripts - Veradigm Sandbox endpoint definition
        // let endpointDefinition = await getEndpointDefinition('7682675b-8247-4fda-b2cd-048bfeafc8af')
        // let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        let authorizeData = await generateFastenConnectAuthorizeUrl(
            'df55610d-8694-4f77-a1d4-3fa45d90a453',
            'dab057b5-e6cc-4369-87ef-3f8bc0d2543c',
            '7682675b-8247-4fda-b2cd-048bfeafc8af',
        )

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
