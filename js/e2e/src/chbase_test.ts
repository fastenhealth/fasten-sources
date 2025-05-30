import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import {generateFastenConnectAuthorizeUrl, generateSourceAuthorizeUrl} from '@shared-library';
import process from 'process';

test("CHbase Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);
        //get the VAHealth Sandbox endpoint definition
        // let endpointDefinition = await getEndpointDefinition('ee5e19b6-4539-4e46-baab-b892061fe448')
        // let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        let authorizeData = await generateFastenConnectAuthorizeUrl(
            '3a3e5a4c-7e7b-48ba-8b13-fce0347a6ce8',
            'f45e5c0f-16af-4c09-8ee4-11ddbb58edba',
            'ee5e19b6-4539-4e46-baab-b892061fe448',
        )

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // // We are on landing page
        await page.waitForSelector("text=Sign in to your account");
        await expect(page).toHaveTitle("Sign in to Unify Login");

        // // We are on the ID.me login page
        // await page.waitForSelector("text=Sign in to ID.me");
        await page.click("label[for='username']", { force: true });
        await page.keyboard.type(process.env.PW_CHBASE_USERNAME);
        await page.click("label[for='password']", { force: true });
        await page.keyboard.type(process.env.PW_CHBASE_PASSWORD);
        await page.click("input[type='submit']");

        await page.waitForSelector('#RecordFacilityList');

        // Select the option by its value
        await page.selectOption('#RecordFacilityList', 'ae1dc61d-e73d-4d85-a273-2f1b0f22b24c');
        await page.click("button[type='submit']");


        await page.waitForSelector("text=Choose what data to share");
        await page.click("input[type='submit']");

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

        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
