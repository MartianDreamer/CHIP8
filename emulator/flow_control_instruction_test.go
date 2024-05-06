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

func Test_jump_to_nnn(t *testing.T) {
	em := Make_chip8(500)
	em.execute_instruction([2]byte{0x1a, 0xac})
	if em.pc != 0x0aac {
		t.Failed()
	}
}

func Test_call_subroutine_nnn(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x1a2b
	em.execute_instruction([2]byte{0x2f, 0x1a})
	if em.sp != __STACK_POS {
		t.Failed()
	}
	if em.sp != __STACK_POS+2 || em.pc != 0x0f1a || em.mem[__STACK_POS] != 0x1a || em.mem[__STACK_POS+1] != 0x2b {
		t.Failed()
	}
}

func Test_skip_if_vx_eq_kk(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x0300
	em.v[0xa] = 0xfa
	em.v[0x1] = 0x11
	em.execute_instruction([2]byte{0x3a, 0xfa})
	if em.pc != 0x0302 {
		t.Failed()
	}
	em.execute_instruction([2]byte{0x31, 0xfa})
	if em.pc != 0x0302 {
		t.Failed()
	}
}

func Test_skip_if_vx_neq_kk(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x0300
	em.v[0xa] = 0xfa
	em.v[0x1] = 0x11
	em.execute_instruction([2]byte{0x3a, 0xfa})
	if em.pc != 0x0300 {
		t.Failed()
	}
	em.execute_instruction([2]byte{0x31, 0xfa})
	if em.pc != 0x0302 {
		t.Failed()
	}
}