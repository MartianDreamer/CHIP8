package emulator

import (
	"fmt"
)

func (emulator *Chip8) execute_alu_instruction(opcode byte, instruction [2]byte) {
	switch opcode {
	case 0x6:
		emulator.exec_set_vx_kk(instruction)
	case 0x7:
		emulator.exec_set_vx_vx_add_kk(instruction)
	case 0x8:
		emulator.exec_opcode_8_ins(instruction)
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
	case 0x5:
	case 0x6:
	case 0x7:
	case 0xe:
	}
}
