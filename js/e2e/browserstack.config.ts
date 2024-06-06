import * as cp from 'child_process';
const clientPlaywrightVersion = cp
  .execSync('npx playwright --version')
  .toString()
  .trim()
  .split(' ')[1];
import * as BrowserStackLocal from'browserstack-local';
import * as util from'util';
import process from 'process';

// BrowserStack Specific Capabilities.
// Set 'browserstack.local:true For Local testing
const caps = {
  browser: 'chrome',
  os: 'osx',
  os_version: 'catalina',
  project: 'fastenhealth/fasten-sources',
  name: 'Fasten Sources e2e Tests',
  build: process.env.GITHUB_REF_NAME || 'local-build',
  buildTag: process.env.GITHUB_SHA || 'local-build-tag',
  'browserstack.username': process.env.BROWSERSTACK_USERNAME || 'USERNAME',
  'browserstack.accessKey': process.env.BROWSERSTACK_ACCESS_KEY || 'ACCESSKEY',
  'browserstack.local': process.env.BROWSERSTACK_LOCAL || true,
  'client.playwrightVersion': clientPlaywrightVersion,
};

export const bsLocal = new BrowserStackLocal.Local();

// replace YOUR_ACCESS_KEY with your key. You can also set an environment variable - "BROWSERSTACK_ACCESS_KEY".
export const BS_LOCAL_ARGS = {
  key: process.env.BROWSERSTACK_ACCESS_KEY || 'ACCESSKEY',
};

// Patching the capabilities dynamically according to the project name.
const patchCaps = (name, title) => {
  let combination = name.split(/@browserstack/)[0];
  let [browerCaps, osCaps] = combination.split(/:/);
  let [browser, browser_version] = browerCaps.split(/@/);
  let osCapsSplit = osCaps.split(/ /);
  let os = osCapsSplit.shift();
  let os_version = osCapsSplit.join(' ');
  caps.browser = browser ? browser : 'chrome';
  caps.os_version = browser_version ? browser_version : 'latest';
  caps.os = os ? os : 'osx';
  caps.os_version = os_version ? os_version : 'catalina';
  caps.name = title;
};

export function getCdpEndpoint(name, title){
    patchCaps(name, title)
    const cdpUrl = `wss://cdp.browserstack.com/playwright?caps=${encodeURIComponent(JSON.stringify(caps))}`
    console.log(`--> ${cdpUrl}`)
    return cdpUrl;
}
