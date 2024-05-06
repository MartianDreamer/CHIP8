package emulator

import (
	"fmt"
	"sync"
	"time"
)

const (
	__MEM_SIZE          = 4096
	__REGISTER_QUANTITY = 16
)

type Chip8 struct {
	mem     [__MEM_SIZE]byte
	v       [__REGISTER_QUANTITY]byte
	i       uint16
	pc      uint16
	sp      uint8
	d_timer uint8
	s_timer uint8
	clock   uint16
}

func Make_chip8(clockSpeed uint16) *Chip8 {
	rs := &Chip8{
		clock: clockSpeed,
		sp: __STACK_POS,
	}
	copy(rs.mem[:], font_sprites[:])
	return rs
}

func (emulator *Chip8) Start() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		emulator.cycle()
	}()
	wg.Wait()
}

func (emulator *Chip8) cycle() {
	var sleepDuration int64 = int64(1000 / emulator.clock)
	start := time.Now().UnixMilli()
	for emulator.pc < __MEM_SIZE {
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
	var ins [2]byte = [2]byte(emulator.mem[emulator.pc : emulator.pc+2])
	emulator.execute_instruction(ins)
	emulator.pc += 2
}

func (emulator *Chip8) execute_instruction(instruction [2]byte) {
	switch opcode := instruction[0] >> 4; opcode {
	case 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x9, 0xb, 0xe:
		emulator.execute_flow_control_instruction(opcode, instruction)
	case 0x6, 0x7, 0x8, 0xa, 0xc, 0xf:
		emulator.execute_alu_instruction(opcode, instruction)
	case 0xd:
		emulator.execute_display_instruction(instruction)
	default:
		panic("unsupported opcode")
	}
}
