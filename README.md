# roboticPump
Hardware info and programming 

Hardware:

-Peresistatic pump

-Raspberry Pi 4

-Breadboard

-Flyback diodes (to protect against arc currents from voltage spikes from Peresistatic pump because it is an inductive load)

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


