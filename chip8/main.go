package main

import (
	"log"
	"github.com/MartianDreamer/CHIP8/emulator"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	em := emulator.Make_chip8(500)
	em.Start()
	game := &emulator.Chip8_Renderer{Emulator: em}
	ebiten.SetWindowSize(320, 160)
	ebiten.SetWindowTitle("MyChip8")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
