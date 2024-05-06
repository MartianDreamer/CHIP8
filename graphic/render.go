package graphic

import (
	"github.com/MartianDreamer/CHIP8/emulator"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

const __MASK = 0b10000000

var __FG_COLOR = color.White
var __BG_COLOR = color.Black

type Renderer struct {
	Emulator *emulator.Chip8
	started  bool
	keys     []ebiten.Key
}

func (r *Renderer) Update() error {
	if !r.started {
		r.Emulator.Start()
		r.started = true
	}
	r.Emulator.Keyboard[0] = 0x0
	r.keys = inpututil.AppendPressedKeys(r.keys[:0])
	if len(r.keys) > 0 {
		clicked, key := keyMap(r.keys[0])
		r.Emulator.Keyboard[0] = clicked
		r.Emulator.Keyboard[1] = key
	}
	return nil
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	image := r.mapToImage()
	screen.DrawImage(image, &ebiten.DrawImageOptions{})
}

func (r *Renderer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}

func (r *Renderer) mapToImage() *ebiten.Image {
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

func keyMap(key ebiten.Key) (byte, byte) {

	switch key {
	case ebiten.Key0:
		return 0x1, 0x0
	case ebiten.Key1:
		return 0x1, 0x1
	case ebiten.Key2:
		return 0x1, 0x2
	case ebiten.Key3:
		return 0x1, 0x3
	case ebiten.Key4:
		return 0x1, 0x4
	case ebiten.Key5:
		return 0x1, 0x5
	case ebiten.Key6:
		return 0x1, 0x6
	case ebiten.Key7:
		return 0x1, 0x7
	case ebiten.Key8:
		return 0x1, 0x8
	case ebiten.Key9:
		return 0x1, 0x9
	case ebiten.KeyQ:
		return 0x1, 0xa
	case ebiten.KeyW:
		return 0x1, 0xb
	case ebiten.KeyE:
		return 0x1, 0xc
	case ebiten.KeyR:
		return 0x1, 0xd
	case ebiten.KeyT:
		return 0x1, 0xe
	case ebiten.KeyY:
		return 0x1, 0xf
	default:
		return 0x0, 0x0
	}
}
