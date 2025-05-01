import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import {generateFastenConnectAuthorizeUrl, generateSourceAuthorizeUrl} from '@shared-library';

test("Epic Non Legacy Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);
        //get the Epic Sandbox endpoint definition
        // let endpointDefinition = await getEndpointDefinition('8e2f5de7-46ac-4067-96ba-5e3f60ad52a4')
        // let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        let authorizeData = await generateFastenConnectAuthorizeUrl(
            'e16b9952-8885-4905-b2e3-b0f04746ed5c',
            '2727ec27-67e9-475a-bea1-423102beaa1d',
            '8e2f5de7-46ac-4067-96ba-5e3f60ad52a4',
        )


        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on MyChart login page
        await page.waitForSelector("text=MyChart Username");
        await expect(page).toHaveTitle("MyChart - Login Page");
        await page.click("label[for='Login']", { force: true });
        await page.keyboard.type("fhircamila");
        await page.click("label[for='Password']", { force: true });
        await page.keyboard.type("epicepic1");
        await page.click("text=Log In");

        // // We have logged in to MyChart
        await page.waitForSelector("text=Fasten Health has said that it:");
        await page.locator('text=Continue') //wait for continue button
        await expect(page).toHaveTitle("MyChart - Are you sure?");
        // await page.getByTitle("Continue to next page").click({ delay: 1000 });
        await page.waitForSelector('button:text("Continue")', { state: 'visible' });
        await page.waitForTimeout(1000)
        await page.getByRole('button', { name: 'Continue' }).click();


        // We are on the MyChart authorize page. Authorize our app for 1 hour.
        await page.waitForSelector("text=What would you like to share?");
        await expect(page).toHaveTitle("MyChart - Are you sure?");
        await page.click('text=3 months', { force: true, delay: 1000 });
        await page.click("text=Allow access", { force: true, delay: 500 });

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
