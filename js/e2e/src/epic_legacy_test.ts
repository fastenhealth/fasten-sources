import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';

test("Epic OAuth2 Legacy Login Flow", async ({page}) => {
    test.skip()

    //get the Epic Sandbox endpoint definition
    let endpointDefinition = await getEndpointDefinition('fc94bfc7-684d-4e4d-aa6e-ceec01c21c81')
    let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

    // authorizeData.sourceState
    console.log(authorizeData.url.toString())

    // Start login flow by clicking on button with text "Login to MyChart"
    await page.goto(authorizeData.url.toString());

    // We are on MyChart login page
    await page.waitForSelector("text=MyChart Username");
    await expect(page).toHaveTitle("MyChart - Login Page");
    await page.click("label[for='Login']", { force: true });
    await page.keyboard.type("fhirderrick");
    await page.click("label[for='Password']", { force: true });
    await page.keyboard.type("epicepic1");
    await page.click("text=Sign In");

    // We have logged in to MyChart
    await page.waitForSelector("text=Fasten Health has said that it:");
    await page.locator('text=Continue') //wait for continue button
    await expect(page).toHaveTitle("MyChart - Are you sure?");
    await page.getByTitle("Continue to next page").click();


    // We are on the MyChart authorize page. Authorize our app for 1 hour.
    await page.waitForSelector("text=What would you like to share?");
    await expect(page).toHaveTitle("MyChart - Are you sure?");
    // await page.click('text=3 months', { force: true, delay: 1000 });
    await page.click("text=Allow access", { force: true, delay: 500 });

    // MyChart has granted access, redirecting back to app from MyChart
    await page.waitForSelector("text=Epic FHIR Dynamic Registration Redirect");

    // Should auto redirect if successful, but playwright Chrome seems to have issues so we'll manually click if needed
    try {
        await page.click("text=Back to main page", { force: true, delay: 5000 });
    } catch (e) {}

    // If successful, Dynamic Client Registration Data should now be visible
    await page.waitForSelector("text=Dynamic Client Registration Data");



    // await expect(page).toHaveTitle("MyChart - Are you sure?");
    // await page.getByTitle("Continue to next page").click({
    //     force: true,
    //     delay: 1000,
    // });
    //
    // // We are on the MyChart authorize page. Authorize our app for 1 hour.
    // await page.waitForSelector("text=What would you like to share?");
    // await expect(page).toHaveTitle("MyChart - Are you sure?");
    // // await page.click('text=3 months', { force: true, delay: 1000 });
    // await page.click("text=Allow access", { force: true, delay: 500 });
    //
    // // Should auto redirect if successful, but playwright Chrome seems to have issues so we'll manually click if needed
    // try {
    //     await page.click("text=Back to main page", { force: true, delay: 5000 });
    // } catch (e) {}
    //
    // // MyChart has granted access, redirecting back to app from MyChart
    // await page.waitForSelector("text=Your account has been securely connected to FASTEN.")
    // await expect((new URL(page.url())).searchParams.get("code")).toBeTruthy()
});
//
// const expirationOptions = ["1 hour", "1 day", "1 week", "1 month", "3 months"];
// for (const timeOption of expirationOptions) {
//     test(`Epic OAuth2 Login Flow with Dynamic Registration Works with '${timeOption}' Expiration Option Selected`, async ({
//                                                                                                                               page,
//                                                                                                                           }) => {
//         await page.goto("http://localhost:5173/");
//
//         // Before we log in, there should be no dynamic client data
//         await page.waitForSelector(
//             "text=After you log in, your dynamic client data will show here"
//         );
//
//         // Start login flow by clicking on button with text "Login to MyChart"
//         await page.click("text=Login to MyChart");
//
//         // We are on MyChart login page
//         await page.waitForSelector("text=MyChart Username");
//         await expect(page).toHaveTitle("MyChart - Login Page");
//         await page.click("label[for='Login']", { force: true });
//         await page.keyboard.type("fhirderrick");
//         await page.click("label[for='Password']", { force: true });
//         await page.keyboard.type("epicepic1");
//         await page.click("text=Sign In");
//
//         // We have logged in to MyChart
//         await page.waitForSelector("text=Not a Company at all has said that it:");
//         await expect(page).toHaveTitle("MyChart - Are you sure?");
//         await page.getByTitle("Continue to next page").click({
//             force: true,
//             delay: 1000,
//         });
//
//         // We are on the MyChart authorize page. Authorize our app for 1 hour.
//         await page.waitForSelector("text=What would you like to share?");
//         await expect(page).toHaveTitle("MyChart - Are you sure?");
//         await page.click(`text=${timeOption}`, { force: true, delay: 1000 });
//         await page.click("text=Allow access", { force: true, delay: 500 });
//
//         // MyChart has granted access, redirecting back to app from MyChart
//         await page.waitForSelector("text=Epic FHIR Dynamic Registration Redirect");
//
//         // Should auto redirect if successful, but playwright Chrome seems to have issues so we'll manually click if needed
//         try {
//             await page.click("text=Back to main page", { force: true, delay: 5000 });
//         } catch (e) {}
//
//         // If successful, Dynamic Client Registration Data should now be visible
//         await page.waitForSelector("text=Dynamic Client Registration Data");
//     });
// }
