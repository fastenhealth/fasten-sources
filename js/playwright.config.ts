// @ts-check
import { devices } from '@playwright/test';
import { getCdpEndpoint } from './e2e/browserstack.config.ts'
import process from 'process';
import * as path from 'path';

/**
 * Read environment variables from file.
 * https://github.com/motdotla/dotenv
 */
// require('dotenv').config();


/**
 * @see https://playwright.dev/docs/test-configuration
 * @type {import('@playwright/test').PlaywrightTestConfig}
 */
export default {
  testDir: './e2e/src',
  testMatch: '**/*.ts',

  globalSetup: path.join(path.dirname("."), 'e2e/global-setup.ts'),
  globalTeardown: path.join(path.dirname("."), './e2e/global-teardown.ts'),

  /* Maximum time one test can run for. */
  timeout: 90 * 1000,
  expect: {
    /**
     * Maximum time expect() should wait for the condition to be met.
     * For example in `await expect(locator).toHaveText();`
     */
    timeout: 10000
  },
  /* Run tests in files in parallel */
  fullyParallel: true,
  /* Fail the build on CI if you accidentally left test.only in the source code. */
  forbidOnly: !!process.env.CI,
  /* Retry on CI only */
  retries: process.env.CI ? 2 : 0,
  /* Opt out of parallel tests on CI. */
  workers: process.env.CI ? 1 : undefined,
  /* Reporter to use. See https://playwright.dev/docs/test-reporters */
  reporter: process.env.CI ? 'html' : 'line',
  /* Shared settings for all the projects below. See https://playwright.dev/docs/api/class-testoptions. */
  use: {
    /* Maximum time each action such as `click()` can take. Defaults to 0 (no limit). */
    actionTimeout: 0,
    /* Base URL to use in actions like `await page.goto('/')`. */
    // baseURL: 'http://localhost:3000',

    /* Collect trace when retrying the failed test. See https://playwright.dev/docs/trace-viewer */
    trace: 'on-first-retry',
  },

  /* Configure projects for major browsers */
  projects: [
    {
      name: 'aetna',
      testMatch: /.*aetna_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'aetna') },
      },
    },
    {
      name: 'allscripts',
      testMatch: /.*allscripts_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'allscripts') },
      },
    },
    {
      name: 'athena',
      testMatch: /.*athena_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'athena') },
      },
    },
    {
      name: 'careevolution',
      testMatch: /.*careevolution_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'careevolution') },
      },
    },
    {
      name: 'cerner',
      testMatch: /.*cerner_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'cerner') },
      },
    },
    {
      name: 'cigna',
      testMatch: /.*cigna_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'cigna') },
      },
    },
    {
      name: 'drchrono',
      testMatch: /.*drchrono_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'drchrono') },
      },
    },
    {
      name: 'dynamichealthit',
      testMatch: /.*dynamichealthit_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'dynamichealthit') },
      },
    },
    {
      name: 'eclinicalworks',
      testMatch: /.*eclinicalworks_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'eclinicalworks') },
      },
    },
    {
      name: 'epic-legacy',
      testMatch: /.*epic_legacy_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'epic-legacy') },
      },
    },
    {
      name: 'humana',
      testMatch: /.*humana_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'humana') },
      },
    },
    {
      name: 'kaiser',
      testMatch: /.*kaiser_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'kaiser') },
      },
    },
    {
      name: 'maximeyes',
      testMatch: /.*maximeyes_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'maximeyes') },
      },
    },
    {
      name: 'medhost',
      testMatch: /.*medhost_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'medhost') },
      },
    },
    {
      name: 'medicare',
      testMatch: /.*medicare_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'medicare') },
      },
    },
    {
      name: 'nextgen',
      testMatch: /.*nextgen_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'nextgen') },
      },
    },
    {
      name: 'onemedical',
      testMatch: /.*onemedical_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'onemedical') },
      },
    },
    {
      name: 'practicefusion',
      testMatch: /.*practicefusion_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'practicefusion') },
      },
    },
    {
      name: 'vahealth',
      testMatch: /.*vahealth_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'vahealth') },
      },
    },
    {
      name: 'chbase',
      testMatch: /.*chbase_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'chbase') },
      },
    },
    {
      name: 'epic-non-legacy',
      testMatch: /.*epic_non_legacy_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'epic-non-legacy') },
      },
    },
    {
      name: 'flatiron',
      testMatch: /.*flatiron_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'flatiron') },
      },
    },
    {
      name: 'united-healthcare',
      testMatch: /.*united_healthcare_test.ts/,
      use: {
        connectOptions: { wsEndpoint: getCdpEndpoint('chrome@latest:Windows 11', 'united-healthcare') },
      },
    }

    // {
    //   name: 'playwright-webkit@latest:OSX Ventura',
    //   use: {
    //     connectOptions: { wsEndpoint: getCdpEndpoint('playwright-webkit@latest:OSX Ventura', 'test2') }
    //   },
    // },
    // {
    //   name: 'playwright-firefox:Windows 11',
    //   use: {
    //     connectOptions: { wsEndpoint: getCdpEndpoint('playwright-firefox:Windows 11', 'test3') }
    //   },
    // }
    ,{
      name: 'wip',
      testMatch: /.*practicefusion_test.ts/,
      use: { ...devices['Desktop Chrome'] },
    }
  ],

  /* Folder for test artifacts such as screenshots, videos, traces, etc. */
  // outputDir: 'test-results/',

  /* Run your local dev server before starting the tests */
  // webServer: {
  //   command: 'npm run start',
  //   port: 3000,
  // },
};
