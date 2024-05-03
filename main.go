package main

import "CHIP8/emulation"

func main() {
	emulator := emulation.MakeChip8()
	emulator.Start()
}
