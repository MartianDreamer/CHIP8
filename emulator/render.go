package emulator

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const __MASK = 0b10000000

var __FG_COLOR = color.Black
var __BG_COLOR = color.White

type Chip8_Renderer struct {
	Emulator *Chip8
}

func (r *Chip8_Renderer) Update() error {
	return nil
}

func (r *Chip8_Renderer) Draw(screen *ebiten.Image) {
	image := r.mapToImage()
	screen.DrawImage(image, &ebiten.DrawImageOptions{})
}

func (r *Chip8_Renderer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}

func (r *Chip8_Renderer) mapToImage() *ebiten.Image {
	image := ebiten.NewImage(64, 32)
	image.Fill(__BG_COLOR)
	for i, val := range r.Emulator.Screen {
		y := i / 8
		for j := 0; j < 8; j++ {
			x := (i%8)*8 + j
			pixel := (val << j) & __MASK
			if pixel == __MASK {
				image.Set(x, y, __FG_COLOR)
			}
		}
	}
	return image
}
