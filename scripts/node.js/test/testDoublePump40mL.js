//https://github.com/sarfata/pi-blaster
//Build and install directly from source

// var piblaster = require('pi-blaster.js');

const timeMilliSec = 1000;
const secondsToPumpDoublePump = 50;
// const pulseWidthMin = 0.05; //SERVO SAFETY VALUE
const pulseWidthMin = 0.00;    //LED VALUE TO TURN OFF ALL THE WAY
const pulseWidthMax = 0.35;

function timeout(ms) {
   return new Promise(resolve => setTimeout(resolve,ms));
 }

pump40mLSinglePump()

async function pump40mLSinglePump() {

      // piblaster.setPwm(21, pulseWidthMax);
      // piblaster.setPwm(23, pulseWidthMax);

      await timeout(secondsToPumpDoublePump*timeMilliSec)

      // piblaster.setPwm(21, pulseWidthMin);
      // piblaster.setPwm(23, pulseWidthMin);

}
