package emulator

import (
	"testing"

)

func Test_exec_set_vx_kk(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.exec_set_vx_kk([2]byte{0x6a, 0xff})
	if emulator.v[0xa] != 0xff {
		t.Fatalf("v[a]: %x", emulator.v[0xa])
	}
}

func Test_exec_set_vx_vx_add_kk(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0xa] = 0x0f
	emulator.exec_set_vx_vx_add_kk([2]byte{0x7a, 0x11})
	if emulator.v[0xa] != 0x20 {
		t.Fatalf("v[a]: %x", emulator.v[0xa])
	}
}

func Test_exec_set_addr_nnn(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.exec_set_addr_nnn([2]byte{0xab, 0xfa})
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
}

func Test_set_vx_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x2] = 0xfa
	emulator.exec_opcode_8_ins([2]byte{0x80,0x20})
	if emulator.v[0x0] != emulator.v[0x2] {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}


func Test_set_vx_vx_or_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b11110000
	emulator.v[0x2] = 0b11000011
	emulator.exec_opcode_8_ins([2]byte{0x80,0x21})
	if emulator.v[0x0] != 0b11110011 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}

func Test_set_vx_vx_and_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b11110000
	emulator.v[0x2] = 0b11000011
	emulator.exec_opcode_8_ins([2]byte{0x80,0x22})
	if emulator.v[0x0] != 0b11000000 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}
func Test_set_vx_vx_xor_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b11110000
	emulator.v[0x2] = 0b11000011
	emulator.exec_opcode_8_ins([2]byte{0x80,0x23})
	if emulator.v[0x0] != 0b00110011 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}

func Test_set_vx_vx_add_vy_less_than_0xff(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b00110000
	emulator.v[0x2] = 0b11000011
	emulator.exec_opcode_8_ins([2]byte{0x80,0x24})
	if emulator.v[0x0] != 0b11110011 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}


func Test_set_vx_vx_add_vy_greater_than_0xff(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b01110000
	emulator.v[0x2] = 0b11000011
	emulator.exec_opcode_8_ins([2]byte{0x80,0x24})
	if emulator.v[0x0] != 0b00110011 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %x", emulator.v[0x0])
	}
}


func Test_set_vx_vx_substract_vy_vx_greater_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 195
	emulator.v[0x2] = 112
	emulator.exec_opcode_8_ins([2]byte{0x80,0x25})
	if emulator.v[0x0] != 83 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vx_substract_vy_vx_less_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b01111111
	emulator.v[0x2] = 0b11111111
	emulator.exec_opcode_8_ins([2]byte{0x80,0x25})
	if emulator.v[0x0] != 0b10000000 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_shift_right_vx(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b00000011
	emulator.exec_opcode_8_ins([2]byte{0x80,0x26})
	if emulator.v[0x0] != 0b00000001 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
	emulator.v[0x0] = 0b00000010
	emulator.exec_opcode_8_ins([2]byte{0x80,0x26})
	if emulator.v[0x0] != 0b00000001 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vy_substract_vx_vx_greater_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b10000000
	emulator.v[0x2] = 0b00111100
	emulator.exec_opcode_8_ins([2]byte{0x80,0x27})
	if emulator.v[0x0] != 0b10111100 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vy_substract_vx_vx_less_than_vy(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b01111111
	emulator.v[0x2] = 0b11111111
	emulator.exec_opcode_8_ins([2]byte{0x80,0x27})
	if emulator.v[0x0] != 0b10000000 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}

func Test_set_vx_vx_shift_left(t *testing.T) {
	emulator := Make_chip8(500)
	emulator.v[0x0] = 0b10000011
	emulator.exec_opcode_8_ins([2]byte{0x80,0x2e})
	if emulator.v[0x0] != 0b00000110 || emulator.v[0xf] != 0x1 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
	emulator.v[0x0] = 0b00000111
	emulator.exec_opcode_8_ins([2]byte{0x80,0x2e})
	if emulator.v[0x0] != 0b00001110 || emulator.v[0xf] != 0x0 {
		t.Fatalf("v[0]: %d", emulator.v[0x0])
	}
}