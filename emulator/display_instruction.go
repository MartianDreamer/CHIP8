package emulator

func (emulator *Chip8) execute_display_instruction(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	y := instruction[1] >> 4
	n := instruction[1] & __LOWER_MASK
	pos_x := emulator.v[x] % 64
	pos_y := emulator.v[y] % 32
	shift := pos_x % 8
	if shift == 0 {
		for i := 0; i < int(n); i++ {
			pos_y_mod := (int(pos_y) + i) % 32
			scr_pos := __SCR_POS + (pos_y_mod)*8 + int(pos_x)
			sprite_pos := int(emulator.i) + i
			if emulator.mem[scr_pos]+emulator.mem[sprite_pos] != emulator.mem[scr_pos]^emulator.mem[sprite_pos] {
				emulator.v[0xf] = 0x1
			}
			emulator.mem[scr_pos] ^= emulator.mem[sprite_pos]
		}
	} else {
		start_byte := int(pos_x / 8)
		end_byte := int((start_byte + 1) % 8)
		for i := 0; i < int(n); i++ {
			pos_y_mod := (int(pos_y) + i) % 32
			scr_start_pos := __SCR_POS + (pos_y_mod)*8 + start_byte
			scr_end_pos := __SCR_POS + (pos_y_mod)*8 + end_byte
			sprite_pos := int(emulator.i) + i
			if emulator.mem[scr_start_pos]+emulator.mem[sprite_pos]>>shift != emulator.mem[scr_start_pos]^emulator.mem[sprite_pos]>>shift {
				emulator.v[0xf] = 0x1
			}
			emulator.mem[scr_start_pos] ^= emulator.mem[sprite_pos] >> shift
			if emulator.v[0xf] != 0x1 && (emulator.mem[scr_end_pos]+emulator.mem[sprite_pos]<<(8-shift) != emulator.mem[scr_end_pos]^emulator.mem[sprite_pos]<<(8-shift)) {
				emulator.v[0xf] = 0x1
			}
			emulator.mem[scr_end_pos] ^= emulator.mem[sprite_pos] << (8 - shift)
		}
	}
}
