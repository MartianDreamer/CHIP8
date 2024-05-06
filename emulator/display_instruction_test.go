package emulator

import "testing"

func Test_display_at_8_divisible_pos_x_without_erasing(t *testing.T) {
	em := Make_chip8(500)
	em.i = 0x0
	em.v[0x0] = 56
	em.v[0x1] = 10
	em.execute_instruction([2]byte{0xd0, 0x15})
	if em.mem[__SCR_POS+(10*__WIDTH)+7] != 0xf0 ||
		em.mem[__SCR_POS+(10*__WIDTH+8)+7] != 0x90 ||
		em.mem[__SCR_POS+(10*__WIDTH+16)+7] != 0x90 ||
		em.mem[__SCR_POS+(10*__WIDTH+24)+7] != 0x90 ||
		em.mem[__SCR_POS+(10*__WIDTH+32)+7] != 0xf0 {
		t.Fatal()
	}
}

func Test_display_at_8_divisible_pos_x_with_erasing(t *testing.T) {
	em := Make_chip8(500)
	em.i = 0x210
	em.mem[0x210] = 0b11110000
	em.mem[0x211] = 0b11110011
	em.mem[0x212] = 0b11110011
	em.v[0x0] = 56
	em.v[0x1] = 10
	em.mem[__SCR_POS+(10*__WIDTH)+7] = 0b10000000
	em.execute_instruction([2]byte{0xd0, 0x13})
	if em.mem[__SCR_POS+(10*__WIDTH)+7] != 0b01110000 ||
		em.mem[__SCR_POS+(10*__WIDTH+8)+7] != 0b11110011 ||
		em.mem[__SCR_POS+(10*__WIDTH+16)+7] != 0b11110011 ||
		em.v[0xf] != 1 {
		t.Fatal()
	}
}

func Test_display_at_not_8_divisible_pos_x_without_erasing(t *testing.T) {
	em := Make_chip8(500)
	em.i = 0x210
	em.mem[0x210] = 0b11110100
	em.mem[0x211] = 0b11110011
	em.mem[0x212] = 0b11110011
	em.v[0x0] = 51
	em.v[0x1] = 10
	em.execute_instruction([2]byte{0xd0, 0x13})
	if em.mem[__SCR_POS+(10*__WIDTH)+6] != 0b00011110 ||
		em.mem[__SCR_POS+(10*__WIDTH+8)+6] != 0b00011110 ||
		em.mem[__SCR_POS+(10*__WIDTH+16)+6] != 0b00011110 ||
		em.v[0xf] != 0 {
		t.Fatal()
	}
	if em.mem[__SCR_POS+(10*__WIDTH)+7] != 0b10000000 ||
		em.mem[__SCR_POS+(10*__WIDTH+8)+7] != 0b01100000 ||
		em.mem[__SCR_POS+(10*__WIDTH+16)+7] != 0b01100000 ||
		em.v[0xf] != 0 {
		t.Fatal()
	}
}

func Test_display_at_not_8_divisible_pos_x_with_erasing(t *testing.T) {
	em := Make_chip8(500)
	em.i = 0x210
	em.mem[0x210] = 0b11110100
	em.mem[0x211] = 0b11110011
	em.mem[0x212] = 0b11110011
	em.mem[198] = 0b00010000
	em.v[0x0] = 51
	em.v[0x1] = 10
	em.execute_instruction([2]byte{0xd0, 0x13})
	if em.mem[__SCR_POS+(10*__WIDTH)+6] != 0b00001110 ||
		em.mem[__SCR_POS+(10*__WIDTH+8)+6] != 0b00011110 ||
		em.mem[__SCR_POS+(10*__WIDTH+16)+6] != 0b00011110 ||
		em.v[0xf] != 1 {
		t.Fatal()
	}
	if em.mem[__SCR_POS+(10*__WIDTH)+7] != 0b10000000 ||
		em.mem[__SCR_POS+(10*__WIDTH+8)+7] != 0b01100000 ||
		em.mem[__SCR_POS+(10*__WIDTH+16)+7] != 0b01100000 ||
		em.v[0xf] != 1 {
		t.Fatal()
	}
}
