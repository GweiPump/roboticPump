# roboticPump

Hardware info and programming

Demo testing hardware:

<img src="https://github.com/GweiPump/roboticPump/blob/main/images/golangSpeed.png" alt="golangIsFast"/>
<img src="https://github.com/GweiPump/roboticPump/blob/main/images/gpioPins.png" alt="gpioPins"/>
<img src="https://github.com/GweiPump/roboticPump/blob/main/images/hardwareWiring.png" alt="hardwareWiring"/>
<img src="https://github.com/GweiPump/roboticPump/blob/main/images/switchTypes.png" alt="switchTypes"/>


https://www.youtube.com/watch?v=CX-vrFobFG4&ab_channel=Singularity2045

Note: Robotic pump priming takes about 2 seconds.

Tubing length both ends for pumps is about 32 cm.

## Gobot flow pump rate (faster) 

Pump rate is about 40 mL every 45 seconds not primed with a single pump.

Pump rate is about 40 mL every 24 seconds not primed with two pumps at the same time.

## pi-blaster.js pump flow rate (slower)

Pump rate is about 40 mL every 76 seconds not primed with a single pump.

Pump rate is about 40 mL every 50 seconds not primed with two pumps at the same time.

Using GPIO pins 21 and 23.

Wiring info:

Wiring diagram online drawing tool:

https://www.circuit-diagram.org/editor/

300 ohm resistor
https://electronics.stackexchange.com/questions/108113/control-12-solenoids-with-a-raspberry-pi

Wiring (with 300 ohm resistor, Raspberry Pi 4 has 3.3V GPIO while Arduino has 5.0V GPIO):
https://www.officialhrm.com/arduino/arduino-tip120-motor

Golang (go1.17 is working with Geth and Gobot):

Install Golang and Golang Version Manager:

https://github.com/moovweb/gvm

Using gobot:

https://gobot.io/documentation/platforms/raspi/

Hardware:

-Peristaltic pump

-Raspberry Pi 4

-Breadboard

-Flyback diodes (to protect against arc currents from voltage spikes from peresistatic pump because it is an inductive load)

-Power supply

-TIP120 (BJT) (to use Raspberry Pi 4 GPIO output pins to control peresistatic pump)

Note 1: you can use a mechanical relay board instead of a transistor (MOSFET [voltage controlled] or BJT [current controlled]).

Transistors:

-Last longer than mechanical relay boards

-Can switch on and off faster (higher frequency) than mechanical relay boards (electricity is faster than moving parts)

Mechanical relay boards:

-Accurate for keeping voltages the same between connections (useful for reading electric sensor data, not needed for this application)

Note 2: Raspberry Pi and transistor comparison:

BJTs:

-Easy to control with Raspberry Pi

-Good for low current applications

MOSFETS:

-Hard to control with Raspberry Pi 4 (3.3V GPIO) instead of Ardunio (5.0V GPIO)

-Have a higher switch frequency than BJTs

-Very energy effecient and easy to control with gate voltages
