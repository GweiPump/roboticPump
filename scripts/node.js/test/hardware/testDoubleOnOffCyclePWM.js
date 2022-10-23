//https://github.com/sarfata/pi-blaster
//Build and install directly from source

// var piblaster = require('pi-blaster.js');
const halfCycle = 2000;
const fullCycle = 2*halfCycle;

// const pulseWidthMin = 0.05; //SERVO SAFETY VALUE
const pulseWidthMin = 0.00;    //LED VALUE TO TURN OFF ALL THE WAY
const pulseWidthMax = 0.35;

function timeout(ms) {
   return new Promise(resolve => setTimeout(resolve,ms));
 }

testDoubleOnOffCyclePWM()

async function testDoubleOnOffCyclePWM() {

    // piblaster.setPwm(21, pulseWidth);
    // piblaster.setPwm(23, pulseWidth);
    console.log("ON")
    await timeout(halfCycle)
    
    // piblaster.setPwm(21, pulseWidth);
    // piblaster.setPwm(23, pulseWidth);
    console.log("OFF")
    await timeout(halfCycle)

    testDoubleOnOffCyclePWM()     //Call function again to loop

}
