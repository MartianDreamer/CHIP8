package emulator

const __LOWER_MASK uint8 = 0b00001111
const __LEAST_SIGNIFICANT_MASK uint8 = 0b00000001
const __FONT_CHARACTER_SIZE = 5
const __FONT_CHARACTER_COUNT = 16
const __STACK_ENTRY_SIZE = 2
const __STACK_ENTRY_COUNT = 16
const __WIDTH = 8
const __HEIGHT = 32
const __FONT_SPRITE_POS = 0
const __STACK_POS = __FONT_SPRITE_POS + (__FONT_CHARACTER_SIZE * __FONT_CHARACTER_COUNT)
const __SCR_POS = __STACK_POS + (__STACK_ENTRY_SIZE * __STACK_ENTRY_COUNT)
const __KB_POS = __SCR_POS + (__WIDTH * __HEIGHT)

var font_sprites = [80]byte{
	0xf0, 0x90, 0x90, 0x90, 0xf0,
	0x20, 0x60, 0x20, 0x20, 0x70,
	0xF0, 0x10, 0xF0, 0x80, 0xF0,
	0xF0, 0x10, 0xF0, 0x10, 0xF0,
	0x90, 0x90, 0xF0, 0x10, 0x10,
	0xF0, 0x80, 0xF0, 0x10, 0xF0,
	0xF0, 0x80, 0xF0, 0x90, 0xF0,
	0xF0, 0x10, 0x20, 0x40, 0x40,
	0xF0, 0x90, 0xF0, 0x90, 0xF0,
	0xF0, 0x90, 0xF0, 0x10, 0xF0,
	0xF0, 0x90, 0xF0, 0x90, 0x90,
	0xE0, 0x90, 0xE0, 0x90, 0xE0,
	0xF0, 0x80, 0x80, 0x80, 0xF0,
	0xE0, 0x90, 0x90, 0x90, 0xE0,
	0xF0, 0x80, 0xF0, 0x80, 0xF0,
	0xF0, 0x80, 0xF0, 0x80, 0x80,
}
