//https://github.com/sarfata/pi-blaster
//Build and install directly from source
var piblaster = require('pi-blaster.js');
const timeMilliSec = 2000;
// const pulseWidthMin = 0.05; //SERVO SAFETY VALUE
const pulseWidthMin = 0.00;    //LED VALUE TO TURN OFF ALL THE WAY
const pulseWidthMax = 0.35;
let pulseWidth = pulseWidthMax

setInterval(() => { 
  piblaster.setPwm(21, pulseWidth); 
  console.log(pulseWidth)
  if (pulseWidth == pulseWidthMin) {
      pulseWidth = pulseWidthMax
  } else if (pulseWidth == pulseWidthMax) {
    	pulseWidth = pulseWidthMin;
  }
}, timeMilliSec);
