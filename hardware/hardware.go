package hardware

type Chip8 struct {
	MEM    [4096]byte
	V      [16]byte
	I      uint16
	PC     uint16
	SP     uint8
	DTIMER uint8
	STIMER uint8
	SCR    [32][8]byte
}

var Emulator = makeChip8();

func makeChip8() *Chip8 {
	rs := &Chip8{}
	preload(rs, fontSprites);
	return rs
}

func preload(emulator *Chip8, sprites [16][5]byte) {
	for i, sprite := range sprites {
		for j := 0; j < 5; j++ {
			emulator.MEM[i*5+j] = sprite[j]
		}
	}
}
