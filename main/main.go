package main

import (
	"log"
	"os"

	"github.com/MartianDreamer/CHIP8/emulator"
	"github.com/MartianDreamer/CHIP8/graphic"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing argument")
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("can't open file")
	}
	em := emulator.Make_chip8(1000)
	em.LoadRom(file)
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("MyChip8")
	if err := ebiten.RunGame(graphic.Make_Renderer(em)); err != nil {
		log.Fatal(err)
	}
}
