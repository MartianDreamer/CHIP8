package emulation

import (
	"fmt"
	"time"
)

type Chip8 struct {
	MEM     [4096]byte
	V       [16]byte
	I       uint16
	PC      uint16
	SP      uint8
	D_TIMER uint8
	S_TIMER uint8
	SCR     [32][8]byte
	CLOCK   uint16
	CUR_INS [2]byte
}

func MakeChip8(clockSpeed uint16) *Chip8 {
	rs := &Chip8{
		CLOCK: clockSpeed,
	}
	copy(fontSprites[:], rs.MEM[:])
	return rs
}

func (emulator *Chip8) Start() {
	var sleepDuration int64 = int64(60000 / emulator.CLOCK)
	start := time.Now().UnixMilli()
	for {
		now := time.Now().UnixMilli()
		if now - start < sleepDuration {
			continue
		}
		start = now
		emulator.fetchInstruction()
		time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	}
}

func (emulator *Chip8) fetchInstruction() {
	fmt.Printf("Fetch instruction at %d\n", emulator.PC)
	emulator.CUR_INS[0] = emulator.MEM[emulator.PC]
	emulator.CUR_INS[1] = emulator.MEM[emulator.PC+1]
	emulator.PC += 2
}
