package emulator

import "time"

const (
	__MEM_SIZE          = 4096
	__REGISTER_QUANTITY = 16
)

type Chip8 struct {
	mem      [__MEM_SIZE]byte
	v        [__REGISTER_QUANTITY]byte
	i        uint16
	pc       uint16
	sp       uint8
	d_timer  uint8
	S_timer  uint8
	clock    uint32
	running  bool
	Screen   []byte
	Keyboard []byte
}

func Make_chip8(clockSpeed uint32) *Chip8 {
	rs := &Chip8{
		clock: clockSpeed,
		sp:    __STACK_POS - 2, // - 2 because there is no stack frame in the stack when we init the emulator
		pc:    __PROGRAM_POS,
	}
	rs.Screen = rs.mem[__SCR_POS : __SCR_POS+(__WIDTH*__HEIGHT)]
	rs.Keyboard = rs.mem[__KB_POS : __KB_POS+2]
	copy(rs.mem[:], font_sprites[:])
	return rs
}

func (emulator *Chip8) Start() {
	go func() {
		emulator.cycle()
	}()
	go func() {
		emulator.timer_cycle()
	}()
	emulator.running = true
}

func (em *Chip8) Stop() {
	em.running = false
}

func (em *Chip8) Reset() {
	em.pc = __PROGRAM_POS
	em.sp = __STACK_POS - 2
}

func (em *Chip8) LoadRom(file []byte) {
	copy(em.mem[__PROGRAM_POS:], file)
}

func (emulator *Chip8) cycle() {
	var sleep_duration int64 = int64(1000 / emulator.clock)
	start := time.Now().UnixMilli()
	for emulator.running && emulator.pc < __MEM_SIZE {
		now := time.Now().UnixMilli()
		if now-start < sleep_duration {
			continue
		}
		instruction := emulator.fetch_instruction()
		emulator.execute_instruction(instruction)
		start = now
	}
	emulator.running = false
}

func (emulator *Chip8) fetch_instruction() [2]byte {
	var ins [2]byte = [2]byte(emulator.mem[emulator.pc : emulator.pc+2])
	emulator.pc += 2
	return ins
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

func (em *Chip8) timer_cycle() {
	sleep_duration := int64(10000 / 60)
	start := time.Now().UnixMilli()
	for em.running {
		now := time.Now().UnixMilli()
		if now-start < sleep_duration {
			continue
		}
		if em.d_timer > 0 {
			em.d_timer--
		}
		if em.S_timer > 0 {
			em.S_timer--
		}
		start = now
	}
}
