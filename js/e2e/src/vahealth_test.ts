import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';

test.skip("VAHealth Login Flow", async ({page}, testInfo) => {
    try {
        await page.evaluate(_ => {},`browserstack_executor: ${JSON.stringify({action: "setSessionName", arguments: {name:testInfo.title}})}`);
        await page.waitForTimeout(5000);
        //get the VAHealth Sandbox endpoint definition
        let endpointDefinition = await getEndpointDefinition('71fd52e3-b9fe-4e3d-b983-99711e798bd8')
        let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        // We are on landing page
        await page.waitForSelector("text=Sign in to VA.gov");
        await expect(page).toHaveTitle("Choose Login Provider");
        await page.click("a.idme-signin", { force: true });

        // We are on the ID.me login page
        await page.waitForSelector("text=Sign in to ID.me");
        await page.click("label[for='user_email']", { force: true });
        await page.keyboard.type("va.api.user+001-2024@gmail.com");
        await page.click("label[for='user_password']", { force: true });
        await page.keyboard.type("SandboxPassword2024!");
        await page.click("input[type='submit']");

        // We are on the ID.me 2FA page
        await page.waitForSelector("text=Complete your sign in");
        await page.click("button[type='submit']");

        // We are on the ID.me 2FA Completion Page
        await page.waitForSelector("text=Complete your sign in");
        await page.click("button[type='submit']");

        // We are on the VA.gov login page
        await page.waitForSelector("text=FastenHealthIncKulatunga-1696557147");
        await expect(page).toHaveTitle("Department of Veteran Affairs Evaluation - Sign In");
        await page.click('button[value="Allow Access"]');

        // If successful, Fasten Lighthouse page should now be visible
        await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'passed',reason: 'Authentication Successful'}})}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => {}, `browserstack_executor: ${JSON.stringify({action: 'setSessionStatus',arguments: {status: 'failed',reason: 'Test failed'}})}`);
    }
});
