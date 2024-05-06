package emulator

import "testing"

func Test_modify_screen(t *testing.T) {
	em := Make_chip8(500)
	em.Screen[0] = 0x1
	if em.mem[__SCR_POS] != 0x1 {
		t.Fatal()
	}
}