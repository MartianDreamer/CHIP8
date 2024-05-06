package emulator

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
	case 0x2:
		emulator.exec_call_subroutine(instruction)
	case 0x3:
		emulator.exec_skip_if_vx_eq_kk(instruction)
	case 0x4:
		emulator.exec_skip_if_vx_neq_kk(instruction)
	case 0x5:
		emulator.exec_skip_if_vx_eq_vy(instruction)
	case 0x9:
		emulator.exec_skip_if_vx_neq_vy(instruction)
	case 0xb:
		emulator.exec_jump_v0_add_nnn(instruction)
	case 0xe:
		emulator.exec_opcode_e_ins(instruction)
	}
}

func (emulator *Chip8) exec_clear_display() {
	for i := __SCR_POS; i < __SCR_POS+(__SCREEN_HEIGHT*__SCREEN_WIDTH_BYTE); i++ {
		emulator.mem[i] = 0x00
	}
}

func (emulator *Chip8) exec_return_from_subroutine() {
	if emulator.sp < __STACK_POS {
		panic("no subroutine to return from")
	}
	emulator.pc = uint16(emulator.mem[emulator.sp])<<8 | uint16(emulator.mem[emulator.sp+1])
	emulator.sp -= 2
}

func (emulator *Chip8) exec_jump_to_nnn(instruction [2]byte) {
	emulator.pc = (uint16(instruction[0]&__LOWER_MASK) << 8) | uint16(instruction[1])
}

func (em *Chip8) exec_call_subroutine(instruction [2]byte) {
	if em.sp >= __SCR_POS {
		panic("nested call exceeded")
	}
	em.sp += 2
	em.mem[em.sp] = byte(em.pc >> 8)
	em.mem[em.sp+1] = byte(em.pc)
	em.pc = uint16(instruction[0]&__LOWER_MASK)<<8 | uint16(instruction[1])
}

func (em *Chip8) exec_skip_if_vx_eq_kk(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	if em.v[x] == instruction[1] {
		em.pc += 2
	}
}

func (em *Chip8) exec_skip_if_vx_neq_kk(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	if em.v[x] != instruction[1] {
		em.pc += 2
	}
}

func (em *Chip8) exec_skip_if_vx_eq_vy(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	y := instruction[1] >> 4
	if em.v[x] == em.v[y] {
		em.pc += 2
	}
}

func (em *Chip8) exec_skip_if_vx_neq_vy(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	y := instruction[1] >> 4
	if em.v[x] != em.v[y] {
		em.pc += 2
	}
}

func (em *Chip8) exec_jump_v0_add_nnn(instruction [2]byte) {
	em.pc = (uint16(instruction[0]&__LOWER_MASK) << 8) | uint16(instruction[1]) + uint16(em.v[0])
}

func (em *Chip8) exec_opcode_e_ins(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	switch instruction[1] {
	case 0x9e:
		if em.mem[__KB_POS] == 1 && em.mem[__KB_POS+1] == em.v[x] {
			em.pc += 2
		}
	case 0xa1:
		if em.mem[__KB_POS] != 1 || em.mem[__KB_POS+1] != em.v[x] {
			em.pc += 2
		}
	}
}
