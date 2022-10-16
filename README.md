# roboticPump
Hardware info and programming

Demo testing hardware:

https://www.youtube.com/watch?v=CX-vrFobFG4&ab_channel=Singularity2045

Pump rate is about 40 mL every 76 seconds not primed per pump.

Wiring info (more to update):

300 ohm resistor
https://electronics.stackexchange.com/questions/108113/control-12-solenoids-with-a-raspberry-pi

Wiring (with 300 ohm resistor, Raspberry Pi 4 has 3.3V GPIO while Arduino has 5.0V GPIO):
https://www.officialhrm.com/arduino/arduino-tip120-motor

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
