# A chip8 emulation

I made this to learn how to emulate hardware

technical reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM

Memory layout:

* Font sprites is located at [0, 80)
* Stack is located at [80,112)
* Screen is located at [112,368)
* Keyboard is located at [368,370)