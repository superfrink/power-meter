Power Meter Tick Counter
================

This is a tick counter using Raspberry Pi GPIO pins.


Input
----------------

Reads from GPIO 27 which is hardware pin 13 on a Raspberry Pi B+.  Be sure to check which hardware pin GPIO 27 maps to on your system or to change the GPIO pin.


Wiring schematic
----------------

FIXME


Building
----------------

Run:

`make`

Generated files:

`power-meter.out` - executable for the OS the build is run on.

`power-meter.pi` - executable for the binary for Linux on ARM 7.
