package emulator

import (
	"fmt"
	"math/rand/v2"
)

func (emulator *Chip8) execute_alu_instruction(opcode byte, instruction [2]byte) {
	switch opcode {
	case 0x6:
		emulator.exec_set_vx_kk(instruction)
	case 0x7:
		emulator.exec_set_vx_vx_add_kk(instruction)
	case 0x8:
		emulator.exec_opcode_8_ins(instruction)
	case 0xa:
		emulator.exec_set_addr_nnn(instruction)
	case 0xc:
		random_number := byte(rand.IntN(0x100))
		emulator.exec_set_vx_rand_and_kk(random_number, instruction)
	case 0xf:
		emulator.exec_opcode_f_ins(instruction)
	}
	fmt.Printf("execute alu instruction %v\n", instruction)
}

func (emulator *Chip8) exec_set_vx_kk(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	emulator.v[x] = instruction[1]
}

func (emulator *Chip8) exec_set_vx_vx_add_kk(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	emulator.v[x] += instruction[1]
}

func (emulator *Chip8) exec_opcode_8_ins(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	y := instruction[1] >> 4
	switch opType := instruction[1] & __LOWER_MASK; opType {
	case 0x0:
		emulator.v[x] = emulator.v[y]
	case 0x1:
		emulator.v[x] |= emulator.v[y]
	case 0x2:
		emulator.v[x] &= emulator.v[y]
	case 0x3:
		emulator.v[x] ^= emulator.v[y]
	case 0x4:
		if check := uint16(emulator.v[x]) + uint16(emulator.v[y]); check > 0xff {
			emulator.v[0xf] = 0x1
		} else {
			emulator.v[0xf] = 0x0
		}
		emulator.v[x] = emulator.v[x] + emulator.v[y]
	case 0x5:
		if emulator.v[x] > emulator.v[y] {
			emulator.v[0xf] = 0x1
		} else {
			emulator.v[0xf] = 0x0
		}
		emulator.v[x] -= emulator.v[y]
	case 0x6:
		emulator.v[0xf] = emulator.v[x] & __LEAST_SIGNIFICANT_MASK
		emulator.v[x] >>= 1
	case 0x7:
		if emulator.v[y] > emulator.v[x] {
			emulator.v[0xf] = 0x1
		} else {
			emulator.v[0xf] = 0x0
		}
		emulator.v[x] = emulator.v[y] - emulator.v[x]
	case 0xe:
		emulator.v[0xf] = emulator.v[x] >> 7
		emulator.v[x] <<= 1
	}
}

func (emulator *Chip8) exec_set_addr_nnn(instruction [2]byte) {
	emulator.i = (uint16(instruction[0]&__LOWER_MASK) << 8) | uint16(instruction[1])
}

func (emulator *Chip8) exec_set_vx_rand_and_kk(random byte, instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	emulator.v[x] = random & instruction[1]
}

func (emulator *Chip8) exec_opcode_f_ins(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	switch instruction[1] {
	case 0x07:
		emulator.v[x] = emulator.d_timer
	case 0x0a:
		if emulator.mem[__KB_POS] != 0x1 {
			emulator.pc -= 2
		}
		emulator.v[x] = emulator.mem[__KB_POS+1]
	case 0x15:
		emulator.d_timer = emulator.v[x]
	case 0x18:
		emulator.s_timer = emulator.v[x]
	case 0x1e:
		emulator.i += uint16(emulator.v[x])
	case 0x29:
		emulator.i = uint16(emulator.v[x] % 16 * 5)
	case 0x33:
		emulator.mem[emulator.i] = emulator.v[x] / 100
		emulator.mem[emulator.i+1] = (emulator.v[x] % 100) / 10
		emulator.mem[emulator.i+2] = emulator.v[x] % 10
	case 0x55:
		for i := 0; i <= int(x); i++ {
			emulator.mem[int(emulator.i)+i] = emulator.v[i]
		}
	case 0x65:
		for i := 0; i <= int(x); i++ {
			emulator.v[i] = emulator.mem[int(emulator.i)+i]
		}
	}
}
