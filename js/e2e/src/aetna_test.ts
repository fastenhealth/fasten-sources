import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';

test("Aetna Login Flow", async ({page}) => {
    test.skip()
    //get the Cerner Sandbox endpoint definition
    let endpointDefinition = await getEndpointDefinition('ac8308d1-90de-4994-bb3d-fe404832714c')
    let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

    // authorizeData.sourceState
    console.log(authorizeData.url.toString())

    // Start login flow by clicking on button with text "Login to MyChart"
    await page.goto(authorizeData.url.toString());

    // We are on login page
    await page.waitForSelector("text=Welcome to Aetna");
    // await expect(page).toHaveTitle("Cerner Health - Sign In");
    await page.click("label[for='username']", { force: true });
    await page.keyboard.type("aetnaTestUser3 ");
    await page.click("label[for='password']", { force: true });
    await page.keyboard.type("FHIRdemo2020");
    await page.click("#loginButton");

    // If successful, Fasten Lighthouse page should now be visible
    await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
});
