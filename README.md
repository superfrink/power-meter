Power Meter Pulse Counter
================

A tick counter using Raspberry Pi GPIO pins.

Created to record electricity used by a model EKM-25IDS meter made by EKM Metering.  The meter's pulse output pins behave much like a switch that closes and opens 800 times per kWh used.


Input
----------------

Reads from GPIO 27 which is hardware pin 13 on a Raspberry Pi B+.  Be sure to check which hardware pin GPIO 27 maps to on your system or to change the GPIO pin.

Output
----------------

Read the number of pulses since the program started by connecting to TCP port 9001.

```
pi@raspberrypi ~ $ while [ 0 ] ; do echo | nc localhost 9001 ; echo ; sleep 60 ; done
10794
10801
10807
10814
```

This allows you to read the value from a remote server and use something like RRDtool to graph power usage.

![Example Power Graph](https://raw.githubusercontent.com/superfrink/power-meter/master/doc/example-power-graph.png)


Wiring
----------------

Schematic:
FIXME

1. Supply 3.3 V from a Raspberry Pi's GPIO pins to the meter.
2. Use a 1k Ohm resistor on the input GPIO pin to limit current.
3. Use a 10k Ohm pulldown resistor on the input pin (after the 1k Ohm resistor).
4. Wire the meter to the GPIO input pin.


Building
----------------

Run:

`make`

Generated files:

`power-meter.out` - executable for the OS the build is run on.

`power-meter.pi` - executable for Linux on ARM 6.
