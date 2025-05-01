import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import {generateFastenConnectAuthorizeUrl, generateSourceAuthorizeUrl} from '@shared-library';
import process from 'process';

test("Flatiron Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);
        // let endpointDefinition = await getEndpointDefinition('22d713f7-c8e5-4e5d-98ea-9cdbe9cfe3fb')
        // let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        let authorizeData = await generateFastenConnectAuthorizeUrl(
            '60da8c03-1e00-4cea-a222-f1c9a1f024ef',
            '09241532-3676-4a57-8f41-3aace3f6a626',
            '22d713f7-c8e5-4e5d-98ea-9cdbe9cfe3fb',
        )

        console.log(authorizeData.url.toString());

        await page.goto(authorizeData.url.toString());

        // // Start login process
        await page.waitForSelector("text=Log in");
        await page.click("label[for='Email']", { force: true });
        await page.keyboard.type(process.env.PW_FLATIRON_USERNAME);
        await page.click("label[for='Password']", { force: true });
        await page.keyboard.type(process.env.PW_FLATIRON_PASSWORD);
        await page.click("input[type='submit']");


        await page.waitForSelector("text=This third-party app is requesting access to your information");
        await page.click("button[type='button']");

        await page.waitForSelector("text=Grant Access?");
        // await page.click("div[data-test='atrium-group'] button:nth-of-type(2)");
        await page.click("div[data-test='atrium-group'] button:has-text('Grant Access')");

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
