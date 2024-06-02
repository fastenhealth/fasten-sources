import { test, expect } from "@playwright/test";
import {getEndpointDefinition} from '../utils';
import {generateSourceAuthorizeUrl} from '../../src/connect/authorization-url';

test.skip("eClinicalWorks-Healow Login Flow", async ({page}) => {
    //get the eClinicalWorks Sandbox endpoint definition
    let endpointDefinition = await getEndpointDefinition('f0a8629a-076c-4f78-b41a-7fc6ae81fa4d')
    let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)

    // authorizeData.sourceState
    console.log(authorizeData.url.toString())

    // Start login flow by clicking on button with text "Login to MyChart"
    await page.goto(authorizeData.url.toString());

    // We are on login page
    await page.waitForSelector("text=FHIR R4 Prodtest EMR");
    await expect(page).toHaveTitle("healow - Health and Online Wellness");
    await page.focus("#username");
    await page.keyboard.type("AdultFemaleFHIR");
    await page.focus("#pwd");
    await page.keyboard.type("e@CWFHIR1");
    await page.click("#btnLoginSubmit");

    // We have logged in
    await page.waitForSelector("text=What you need to know about Fasten Health");
    await expect(page).toHaveTitle("LoginUi");
    await page.click('button:text(" Continue ")');


    // We are on the agreement page
    await page.waitForSelector("text=Personal Information Sharing");
    await expect(page).toHaveTitle("LoginUi");
    //this is a case-sensitive, but not exact match. It may break in the future.
    await page.click('button:text-is("Approve")');


    // If successful, Fasten Lighthouse page should now be visible
    await page.waitForSelector("text=Your account has been securely connected to FASTEN.");
});
