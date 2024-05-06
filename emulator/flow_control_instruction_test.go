package emulator

import "testing"

func Test_exec_clear_display(t *testing.T) {
	em := Make_chip8(500)
	display := []byte{0xff, 0xfa, 0x12, 0x1a, 0xa9, 0x08}
	copy(em.mem[__SCR_POS:], display)
	if em.mem[__SCR_POS] != display[0] {
		t.Failed()
	}
	em.execute_instruction([2]byte{0x00, 0xe0})
	if em.mem[__SCR_POS] != 0x00 {
		t.Failed()
	}
}

func Test_return_from_subroutine(t *testing.T) {
	em := Make_chip8(500)
	em.mem[__STACK_POS] = 0xfa
	em.mem[__STACK_POS+1] = 0xbc
	em.mem[__STACK_POS+2] = 0x12
	em.mem[__STACK_POS+3] = 0xfa
	em.sp = __STACK_POS + 2
	em.execute_instruction([2]byte{0x00, 0xee})
	if em.sp != __STACK_POS || em.pc != 0x12fa {
		t.Failed()
	}
	em.execute_instruction([2]byte{0x00, 0xee})
	if em.sp != __STACK_POS || em.pc != 0xfabc {
		t.Failed()
	}
}
