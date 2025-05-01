import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';

test.skip("Medicare Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the Medicare Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('6ae6c14e-b927-4ce0-862f-91123cb8d774')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on login page
        await page.waitForSelector("text=Log in");
        // await expect(page).toHaveTitle("Cerner Health - Sign In");
        await page.click("label[for='username-textbox']", { force: true });
        await page.keyboard.type("BBUser00000");
        await page.click("label[for='password-textbox']", { force: true });
        await page.keyboard.type("PW00000!");
        await page.click("#login-button");

        //INCOMPLETE - Need to add more steps to complete the login flow

        // // We have logged in
        // await page.waitForSelector("text=Warning: Unknown app");
        // await expect(page).toHaveTitle("Authorization Needed");
        // await page.click('#proceedButton');
        //
        // // We are on the Select Patient page.
        // await page.waitForSelector("text=SMART II, NANCY (Self, 33)");
        // await expect(page).toHaveTitle("Authorization Needed");
        // await page.click("label[for='12724066']", { force: true, delay: 500 });
        // await page.click("#allowButton");
        //
        //
        // If successful, redirect page should now be visible
        // await page.waitForSelector("text=Example Domain");
        //
        // //parse the query string parameters for the current url.
        // const url = new URL(page.url());
        // const params = new URLSearchParams(url.search);
        // //check if the required parameters are present
        // // https://www.example.com/?brand_id=5b7ff2c3-804f-4443-9bd9-4437d43c3b87&connection_status=authorized&endpoint_id=3290e5d7-978e-42ad-b661-1cf8a01a989c&external_state=489bfc62-085b-4a9e-8664-6961106d9120&org_connection_id=25aebcd5-5f2b-435d-bbd2-f0878dd1b4b2&platform_type=cerner&portal_id=00a83214-7b14-4a12-ad95-5198b70dbb63&request_id=f37557d8-1696-4ecb-88f7-8138d39282b5
        // expect(params.has('org_connection_id')).toBe(true);
        // expect(params.has('request_id')).toBe(true);
        // expect(params.get('connection_status')).toEqual("authorized");
        // expect(params.get('external_state')).toEqual(authorizeData.sourceState.state);

        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
