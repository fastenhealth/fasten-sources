import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateFastenConnectAuthorizeUrl, generateSourceAuthorizeUrl} from '@shared-library';
import process from 'process';

//TODO: BROKEN - client id unknown
test.skip("DrChrono Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the DrChrono Sandbox endpoint definition
        // let endpointDefinition = await getEndpointDefinition('6a01ffff-d73e-4728-a315-9b23bd77a4cc')
        // let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        let authorizeData = await generateFastenConnectAuthorizeUrl(
            '575644ea-0ea8-47ce-bbad-149f1e8e4d05',
            'c5aa62ac-71e4-47b3-a61a-96117eeb1d9c',
            '6a01ffff-d73e-4728-a315-9b23bd77a4cc',
        )

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Sign In");
        await page.focus("#username");
        await page.keyboard.type(process.env.PW_DRCHRONO_USERNAME);
        await page.focus("#password");
        await page.keyboard.type(process.env.PW_DRCHRONO_PASSWORD);
        await page.click('button:text("Sign In")');

        // We have logged in
        await page.waitForSelector("text=Charlie DeBraga");
        await page.click('button:text("Select")');

        // Keep me signed in.
        await page.waitForSelector("text=Remember My Decision");
        await page.click('button:text("Yes, Allow")');

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
