package main

import em "github.com/MartianDreamer/CHIP8/emulator"

func main() {
	emulator := em.Make_chip8(500)
	emulator.Start()
}
