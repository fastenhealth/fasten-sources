import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import { generateSourceAuthorizeUrl } from '@shared-library';
import process from 'process';

test.skip("UHC Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);
        let endpointDefinition = await getEndpointDefinition('01b93c2d-df44-4fa7-893f-008cf509e40c')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        console.log(authorizeData.url.toString());

        await page.goto(authorizeData.url.toString());

        // // // Start login process
        await page.waitForSelector("text=Sign in");
        await page.click("label[for='lbl_username']", { force: true });
        await page.keyboard.type(process.env.PW_FLATIRON_USERNAME);
        await page.click("#btnLogin");

        await page.waitForSelector("text=Enter Your Password");
        await page.click("label[for='lbl_login-pwd']", { force: true });
        await page.keyboard.type(process.env.PW_FLATIRON_PASSWORD);
        await page.click("#btnLogin");


        await page.waitForSelector("text=Account Locked");
        await page.click("#nextbtn");


        // await page.click("input[type='submit']");


        // await page.waitForSelector("text=This third-party app is requesting access to your information");
        // await page.click("button[type='button']");

        // await page.waitForSelector("text=Grant Access?");
        // // await page.click("div[data-test='atrium-group'] button:nth-of-type(2)");
        // await page.click("div[data-test='atrium-group'] button:has-text('Grant Access')");



        // await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
