package emulator

import "fmt"

const (
	__SCREEN_HEIGHT     = 32
	__SCREEN_WIDTH_BYTE = 64 / 8
)

func (emulator *Chip8) execute_flow_control_instruction(opcode byte, instruction [2]byte) {
	switch opcode {
	case 0x0:
		opType := instruction[1] & __LOWER_MASK
		if opType == 0x0 {
			emulator.exec_clear_display()
		} else if opType == 0xe {
			emulator.exec_return_from_subroutine()
		}
	case 0x1:
		emulator.exec_jump_to_nnn(instruction)
	default:
		fmt.Println("executed")
	}
	fmt.Printf("execute flow control instruction %v\n", instruction)
}

func (emulator *Chip8) exec_clear_display() {
	for i := __SCR_POS; i < __SCR_POS+(__SCREEN_HEIGHT*__SCREEN_WIDTH_BYTE); i++ {
		emulator.mem[i] = 0x00
	}
}

func (emulator *Chip8) exec_return_from_subroutine() {
	emulator.pc = emulator.sp
	emulator.sp--
}

func (emulator *Chip8) exec_jump_to_nnn(instruction [2]byte) {
	emulator.pc = (uint16(instruction[0]&__LOWER_MASK) << 8) | uint16(instruction[1])
}
