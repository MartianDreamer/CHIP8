package emulator

import (
	"fmt"
	"time"
)

type Chip8 struct {
	mem     [4096]byte
	v       [16]byte
	i       uint16
	pc      uint16
	sp      uint8
	d_timer uint8
	s_timer uint8
	scr     [32][8]byte
	clock   uint16
	cur_ins [2]byte
}

func Make_chip8(clockSpeed uint16) *Chip8 {
	rs := &Chip8{
		clock: clockSpeed,
	}
	copy(FONT_SPRITES[:], rs.mem[:])
	return rs
}

func (emulator *Chip8) Start() {
	var sleepDuration int64 = int64(60000 / emulator.clock)
	start := time.Now().UnixMilli()
	for {
		now := time.Now().UnixMilli()
		if now-start < sleepDuration {
			continue
		}
		start = now
		emulator.fetch_instruction()
		time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	}
}

func (emulator *Chip8) fetch_instruction() {
	fmt.Printf("Fetch instruction at %d\n", emulator.pc)
	emulator.cur_ins[0] = emulator.mem[emulator.pc]
	emulator.cur_ins[1] = emulator.mem[emulator.pc+1]
	emulator.pc += 2
}
