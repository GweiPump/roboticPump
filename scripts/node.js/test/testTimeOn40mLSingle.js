//https://github.com/sarfata/pi-blaster
//Build and install directly from source

// var piblaster = require('pi-blaster.js');

const timeMilliSec = 1000;

// const pulseWidthMin = 0.05; //SERVO SAFETY VALUE
const pulseWidthMin = 0.00;    //LED VALUE TO TURN OFF ALL THE WAY
const pulseWidthMax = 0.35;
let pulseWidth = pulseWidthMax

const startTime = Date.now();

setInterval(() => {

    // piblaster.setPwm(21, pulseWidth);

    console.log("Seconds passed: " + (Date.now() - startTime)/1000) //Pumped to 40mL in about 76 seconds (not primed).

}, timeMilliSec);
