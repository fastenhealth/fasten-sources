// global-setup.js
import { bsLocal, BS_LOCAL_ARGS } from './browserstack.config';
import { promisify } from 'util';
const sleep = promisify(setTimeout);
const redColour = '\x1b[31m';
const whiteColour = '\x1b[0m';

export default async function setup() {
  console.log('Starting BrowserStackLocal ...');
  // Starts the Local instance with the required arguments
  let localResponseReceived = false;
  bsLocal.start(BS_LOCAL_ARGS, (err) => {
    if (err) {
      console.error(
        `${redColour}Error starting BrowserStackLocal${whiteColour}`, err
      );
    } else {
      console.log('BrowserStackLocal Started');
    }
    localResponseReceived = true;
  });
  while (!localResponseReceived) {
    await sleep(1000);
  }
};
