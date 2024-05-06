package emulator

import (
	"testing"
)

func Test_exec_clear_display(t *testing.T) {
	em := Make_chip8(500)
	display := []byte{0xff, 0xfa, 0x12, 0x1a, 0xa9, 0x08}
	copy(em.mem[__SCR_POS:], display)
	if em.mem[__SCR_POS] != display[0] {
		t.Fatal()
	}
	em.execute_instruction([2]byte{0x00, 0xe0})
	if em.mem[__SCR_POS] != 0x00 {
		t.Fatal()
	}
}

func Test_return_from_subroutine(t *testing.T) {
	em := Make_chip8(500)
	em.mem[__STACK_POS] = 0xfa
	em.mem[__STACK_POS+1] = 0xbc
	em.sp = __STACK_POS
	em.execute_instruction([2]byte{0x00, 0xee})
	if em.sp != __STACK_POS-2 || em.pc != 0xfabc {
		t.Fatal()
	}
}

func Test_jump_to_nnn(t *testing.T) {
	em := Make_chip8(500)
	em.execute_instruction([2]byte{0x1a, 0xac})
	if em.pc != 0x0aac {
		t.Fatal()
	}
}

func Test_call_subroutine_nnn(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x1a2b
	if em.sp != __STACK_POS-2 {
		t.Fatal()
	}
	em.execute_instruction([2]byte{0x2f, 0x1a})
	if em.sp != __STACK_POS || em.pc != 0x0f1a || em.mem[__STACK_POS] != 0x1a || em.mem[__STACK_POS+1] != 0x2b {
		t.Fatal()
	}
}

func Test_skip_if_vx_eq_kk(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x0300
	em.v[0xa] = 0xfa
	em.v[0x1] = 0x11
	em.execute_instruction([2]byte{0x3a, 0xfa})
	if em.pc != 0x0302 {
		t.Fatal()
	}
	em.execute_instruction([2]byte{0x31, 0xfa})
	if em.pc != 0x0302 {
		t.Fatal()
	}
}

func Test_skip_if_vx_neq_kk(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x0300
	em.v[0xa] = 0xfa
	em.v[0x1] = 0x11
	em.execute_instruction([2]byte{0x4a, 0xfa})
	if em.pc != 0x0300 {
		t.Fatal()
	}
	em.execute_instruction([2]byte{0x41, 0xfa})
	if em.pc != 0x0302 {
		t.Fatal()
	}
}

func Test_skip_if_vx_eq_vy(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x0300
	em.v[0xa] = 0xfa
	em.v[0x1] = 0xfa
	em.execute_instruction([2]byte{0x5a, 0x10})
	if em.pc != 0x0302 {
		t.Fatal()
	}
}

func Test_skip_if_vx_neq_vy(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x0300
	em.v[0xa] = 0xfa
	em.v[0x1] = 0x11
	em.execute_instruction([2]byte{0x9a, 0x10})
	if em.pc != 0x0302 {
		t.Fatal()
	}
}

func Test_jump_v0_add_nnn(t *testing.T) {
	em := Make_chip8(500)
	em.v[0x0] = 0x02
	em.execute_instruction([2]byte{0xba, 0xb2})
	if em.pc != 0xab4 {
		t.Fatal()
	}
}

func Test_skip_if_vx_pressed(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x300
	em.v[0x8] = 0xb
	em.mem[__KB_POS] = 0x0
	em.execute_instruction([2]byte{0xe8, 0x9e})
	if em.pc != 0x300 {
		t.Fatal()
	}
	em.mem[__KB_POS] = 0x1
	em.mem[__KB_POS+1] = 0xa
	em.execute_instruction([2]byte{0xe8, 0x9e})
	if em.pc != 0x300 {
		t.Fatal()
	}
	em.mem[__KB_POS] = 0x1
	em.mem[__KB_POS+1] = 0xb
	em.execute_instruction([2]byte{0xe8, 0x9e})
	if em.pc != 0x302 {
		t.Fatal()
	}
}

func Test_skip_if_vx_not_pressed(t *testing.T) {
	em := Make_chip8(500)
	em.pc = 0x300
	em.v[0x8] = 0xb
	em.mem[__KB_POS] = 0x1
	em.mem[__KB_POS+1] = 0xb
	em.execute_instruction([2]byte{0xe8, 0xa1})
	if em.pc != 0x300 {
		t.Fatal()
	}
	em.mem[__KB_POS] = 0x1
	em.mem[__KB_POS+1] = 0xa
	em.execute_instruction([2]byte{0xe8, 0xa1})
	if em.pc != 0x302 {
		t.Fatal()
	}
	em.mem[__KB_POS] = 0x0
	em.mem[__KB_POS+1] = 0xb
	em.execute_instruction([2]byte{0xe8, 0xa1})
	if em.pc != 0x304 {
		t.Fatal()
	}
}
