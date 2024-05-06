package emulator

import (
	"testing"
)

func Test_exec_set_vx_kk(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.execute_instruction([2]byte{0x6a, 0xff})
	if emulator.v[0xa] != 0xff {
		t.Fatalf("v[a]: %x", emulator.v[0xa])
	}
}

func Test_exec_set_vx_vx_add_kk(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0xa] = 0x0f
	emulator.execute_instruction([2]byte{0x7a, 0x11})
	if emulator.v[0xa] != 0x20 {
		t.Fatalf("v[a]: %x", emulator.v[0xa])
	}
}

func Test_exec_set_addr_nnn(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.execute_instruction([2]byte{0xab, 0xfa})
	if emulator.i != 0xbfa {
		t.Fatalf("i: %x", emulator.i)
	}
}

func Test_exec_set_vx_rand_and_kk(t *testing.T) {
	emulator := Make_chip8(500)
	var random byte = 0b11110000
	emulator.exec_set_vx_rand_and_kk(random, [2]byte{0xce, 0b11001111})
	if emulator.v[0xe] != 0b11000000 {
		t.Fatalf("v[e]: %x", emulator.v[0xe])
	}
	emulator.execute_instruction([2]byte{0xce, 0b11001111})
	first := emulator.v[0xe]
	emulator.execute_instruction([2]byte{0xce, 0b11001111})
	second := emulator.v[0xe]
	if first == second {
		t.Failed()
	}
}

func Test_set_vx_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x2] = 0xfa
	emulator.execute_instruction([2]byte{0x80, 0x20})
	if emulator.v[0x0] != emulator.v[0x2] {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}

func Test_set_vx_vx_or_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b11110000
	emulator.v[0x2] = 0b11000011
	emulator.execute_instruction([2]byte{0x80, 0x21})
	if emulator.v[0x0] != 0b11110011 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}

func Test_set_vx_vx_and_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b11110000
	emulator.v[0x2] = 0b11000011
	emulator.execute_instruction([2]byte{0x80, 0x22})
	if emulator.v[0x0] != 0b11000000 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}
func Test_set_vx_vx_xor_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b11110000
	emulator.v[0x2] = 0b11000011
	emulator.execute_instruction([2]byte{0x80, 0x23})
	if emulator.v[0x0] != 0b00110011 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}

func Test_set_vx_vx_add_vy_less_than_0xff(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b00110000
	emulator.v[0x2] = 0b11000011
	emulator.execute_instruction([2]byte{0x80, 0x24})
	if emulator.v[0x0] != 0b11110011 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}

func Test_set_vx_vx_add_vy_greater_than_0xff(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b01110000
	emulator.v[0x2] = 0b11000011
	emulator.execute_instruction([2]byte{0x80, 0x24})
	if emulator.v[0x0] != 0b00110011 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}

func Test_set_vx_vx_substract_vy_vx_greater_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 195
	emulator.v[0x2] = 112
	emulator.execute_instruction([2]byte{0x80, 0x25})
	if emulator.v[0x0] != 83 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vx_substract_vy_vx_less_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b01111111
	emulator.v[0x2] = 0b11111111
	emulator.execute_instruction([2]byte{0x80, 0x25})
	if emulator.v[0x0] != 0b10000000 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_shift_right_vx(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b00000011
	emulator.execute_instruction([2]byte{0x80, 0x26})
	if emulator.v[0x0] != 0b00000001 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
	emulator.v[0x0] = 0b00000010
	emulator.execute_instruction([2]byte{0x80, 0x26})
	if emulator.v[0x0] != 0b00000001 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vy_substract_vx_vx_greater_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b10000000
	emulator.v[0x2] = 0b00111100
	emulator.execute_instruction([2]byte{0x80, 0x27})
	if emulator.v[0x0] != 0b10111100 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vy_substract_vx_vx_less_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b01111111
	emulator.v[0x2] = 0b11111111
	emulator.execute_instruction([2]byte{0x80, 0x27})
	if emulator.v[0x0] != 0b10000000 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vx_shift_left(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b10000011
	emulator.execute_instruction([2]byte{0x80, 0x2e})
	if emulator.v[0x0] != 0b00000110 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
	emulator.v[0x0] = 0b00000111
	emulator.execute_instruction([2]byte{0x80, 0x2e})
	if emulator.v[0x0] != 0b00001110 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_d_timer(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.d_timer = 0xf1
	emulator.execute_instruction([2]byte{0xf1, 0x07})
	if emulator.v[0x1] != 0xf1 {
		t.Fatalf("v[1]: %d", emulator.v[0x1])
	}
}

func Test_set_d_timer_vx(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x1] = 0xf1
	emulator.execute_instruction([2]byte{0xf1, 0x15})
	if emulator.d_timer != 0xf1 {
		t.Fatalf("delay timer: %d", emulator.d_timer)
	}
}

func Test_set_s_timer_vx(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x1] = 0xf1
	emulator.execute_instruction([2]byte{0xf1, 0x18})
	if emulator.s_timer != 0xf1 {
		t.Fatalf("delay timer: %d", emulator.d_timer)
	}
}

func Test_set_i_i_add_vx(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.i = 100
	emulator.v[0x2] = 200
	emulator.execute_instruction([2]byte{0xf2, 0x1e})
	if emulator.i != 300 {
		t.Fatalf("i: %d", emulator.i)
	}
}

func Test_set_i_font_location(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x2] = 0xd
	emulator.execute_instruction([2]byte{0xf2, 0x29})
	if emulator.i != 65 {
		t.Fatalf("i: %d", emulator.i)
	}
}

func Test_bcd(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x2] = 203
	emulator.i = 0x300
	emulator.execute_instruction([2]byte{0xf2, 0x33})
	if emulator.mem[0x300] != 2 || emulator.mem[0x301] != 0 || emulator.mem[0x302] != 3 {
		t.Fatalf("mem i: %d\nmem i+1: %d\nmem i+2: %d\n", emulator.mem[emulator.i], emulator.mem[emulator.i+1], emulator.mem[emulator.i+2])
	}
}

func Test_store_register(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0xfa
	emulator.v[0x1] = 0x1b
	emulator.v[0x2] = 0xff
	emulator.v[0x3] = 0x25
	emulator.i = 0x200
	emulator.execute_instruction([2]byte{0xf3, 0x55})
	if emulator.mem[0x200] != 0xfa || emulator.mem[0x201] != 0x1b || emulator.mem[0x202] != 0xff || emulator.mem[0x203] != 0x25 {
		t.Failed()
	}
}

func Test_read_register(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.mem[0x300] = 0xfa
	emulator.mem[0x301] = 0x1b
	emulator.mem[0x302] = 0xff
	emulator.mem[0x303] = 0x25
	emulator.i = 300
	emulator.execute_instruction([2]byte{0xf3, 0x65})
	if emulator.v[0] != 0xfa || emulator.v[1] != 0x1b || emulator.v[2] != 0xff || emulator.v[3] != 0x25 {
		t.Failed()
	}
}

func Test_set_vx_key(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.pc = 0x215
	emulator.execute_instruction([2]byte{0xf0, 0x0a})
	if emulator.pc != 0x213 {
		t.Fatalf("pc: %x\n", emulator.pc)
	}
	emulator.mem[__KB_POS] = 0x1
	emulator.mem[__KB_POS+1] = 0x5
	emulator.pc = 0x215
	emulator.execute_instruction([2]byte{0xf0, 0x0a})
	if emulator.pc != 0x215 || emulator.v[0] != 0x5 {
		t.Failed()
	}
}
