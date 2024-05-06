package emulator

func (emulator *Chip8) execute_display_instruction(instruction [2]byte) {
	x := instruction[0] & __LOWER_MASK
	y := instruction[1] >> 4
	n := instruction[1] & __LOWER_MASK
	pos_x := emulator.v[x] % (__WIDTH * 8)
	pos_y := emulator.v[y] % __HEIGHT
	shift := pos_x % 8
	if shift == 0 {
		pos_x_byte := int(pos_x) / 8
		for i := 0; i < int(n); i++ {
			pos_y_mod := (int(pos_y) + i) % __HEIGHT
			scr_pos := __SCR_POS + (pos_y_mod)*__WIDTH + pos_x_byte
			sprite_pos := int(emulator.i) + i
			if emulator.mem[scr_pos]&emulator.mem[sprite_pos] != 0 {
				emulator.v[0xf] = 0x1
			}
			emulator.mem[scr_pos] ^= emulator.mem[sprite_pos]
		}
	} else {
		start_byte := int(pos_x / 8)
		end_byte := int((start_byte + 1) % 8)
		for i := 0; i < int(n); i++ {
			pos_y_mod := (int(pos_y) + i) % __HEIGHT
			scr_start_pos := __SCR_POS + (pos_y_mod)*__WIDTH + start_byte
			scr_end_pos := __SCR_POS + (pos_y_mod)*__WIDTH + end_byte
			sprite_pos := int(emulator.i) + i
			if emulator.mem[scr_start_pos]&(emulator.mem[sprite_pos]>>shift) != 0 {
				emulator.v[0xf] = 0x1
			}
			emulator.mem[scr_start_pos] ^= emulator.mem[sprite_pos] >> shift
			if emulator.v[0xf] != 0x1 && (emulator.mem[scr_end_pos]&(emulator.mem[sprite_pos]<<(8-shift)) != 0) {
				emulator.v[0xf] = 0x1
			}
			emulator.mem[scr_end_pos] ^= emulator.mem[sprite_pos] << (8 - shift)
		}
	}
}
