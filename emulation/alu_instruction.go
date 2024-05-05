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
		emulator.exec_set_vx_rand_and_kk(instruction)
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
		emulator.v[x] = emulator.v[x] + emulator.v[y]
		if check := uint16(emulator.v[x]) + uint16(emulator.v[y]); check > 255 {
			emulator.v[0xf] = 0x1
		} else {
			emulator.v[0xf] = 0x0
		}
	case 0x5:
		emulator.v[x] -= emulator.v[y]
		if emulator.v[x] > emulator.v[y] {
			emulator.v[0xf] = 0x1
		} else {
			emulator.v[0xf] = 0x0
		}
	case 0x6:
		emulator.v[0xf] = emulator.v[x] & __LEAST_SIGNIFICANT_MASK
		emulator.v[x] >>= 1
	case 0x7:
		emulator.v[x] = emulator.v[y] - emulator.v[x]
		if emulator.v[y] > emulator.v[x] {
			emulator.v[0xf] = 0x1
		} else {
			emulator.v[0xf] = 0x0
		}
	case 0xe:
		emulator.v[0xf] = emulator.v[x] >> 7
		emulator.v[x] <<= 1
	}
}

func (emulator *Chip8) exec_set_addr_nnn(instruction [2]byte) {
	emulator.i = (uint16(instruction[0]&__LOWER_MASK) << 8) | uint16(instruction[1])
}

func (emulator *Chip8) exec_set_vx_rand_and_kk(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	random_number := byte(rand.IntN(256))
	emulator.v[x] = random_number & instruction[1]
}
