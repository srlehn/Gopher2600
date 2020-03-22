// generated code - do not change

package instructions

// GetDefinitions returns the table of instruction definitions for the 6507
func GetDefinitions() ([]*Definition, error) {
	return []*Definition{
		&Definition{OpCode: 0x0, Mnemonic: "BRK", Bytes: 1, Cycles: 7, AddressingMode: 0, PageSensitive: false, Effect: 5},
		&Definition{OpCode: 0x1, Mnemonic: "ORA", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0x3, Mnemonic: "slo", Bytes: 2, Cycles: 8, AddressingMode: 6, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0x4, Mnemonic: "nop", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x5, Mnemonic: "ORA", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x6, Mnemonic: "ASL", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0x7, Mnemonic: "slo", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0x8, Mnemonic: "PHP", Bytes: 1, Cycles: 3, AddressingMode: 0, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x9, Mnemonic: "ORA", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa, Mnemonic: "ASL", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0xc, Mnemonic: "skw", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xd, Mnemonic: "ORA", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xe, Mnemonic: "ASL", Bytes: 3, Cycles: 6, AddressingMode: 3, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x10, Mnemonic: "BPL", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0x11, Mnemonic: "ORA", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0x14, Mnemonic: "nop", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x15, Mnemonic: "ORA", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x16, Mnemonic: "ASL", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x18, Mnemonic: "CLC", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x19, Mnemonic: "ORA", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0x1c, Mnemonic: "skw", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x1d, Mnemonic: "ORA", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x1e, Mnemonic: "ASL", Bytes: 3, Cycles: 7, AddressingMode: 8, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x20, Mnemonic: "JSR", Bytes: 3, Cycles: 6, AddressingMode: 3, PageSensitive: false, Effect: 4},
		&Definition{OpCode: 0x21, Mnemonic: "AND", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0x24, Mnemonic: "BIT", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x25, Mnemonic: "AND", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x26, Mnemonic: "ROL", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x28, Mnemonic: "PLP", Bytes: 1, Cycles: 4, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x29, Mnemonic: "AND", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x2a, Mnemonic: "ROL", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x2b, Mnemonic: "anc", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x2c, Mnemonic: "BIT", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x2d, Mnemonic: "AND", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x2e, Mnemonic: "ROL", Bytes: 3, Cycles: 6, AddressingMode: 3, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x30, Mnemonic: "BMI", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0x31, Mnemonic: "AND", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		nil,
		nil,
		nil,
		&Definition{OpCode: 0x35, Mnemonic: "AND", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x36, Mnemonic: "ROL", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0x37, Mnemonic: "rla", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0x38, Mnemonic: "SEC", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x39, Mnemonic: "AND", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0x3c, Mnemonic: "skw", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x3d, Mnemonic: "AND", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x3e, Mnemonic: "ROL", Bytes: 3, Cycles: 7, AddressingMode: 8, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x40, Mnemonic: "RTI", Bytes: 1, Cycles: 6, AddressingMode: 0, PageSensitive: false, Effect: 5},
		&Definition{OpCode: 0x41, Mnemonic: "EOR", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 0},
		nil,
		nil,
		nil,
		&Definition{OpCode: 0x45, Mnemonic: "EOR", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x46, Mnemonic: "LSR", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x48, Mnemonic: "PHA", Bytes: 1, Cycles: 3, AddressingMode: 0, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x49, Mnemonic: "EOR", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x4a, Mnemonic: "LSR", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x4b, Mnemonic: "asr", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x4c, Mnemonic: "JMP", Bytes: 3, Cycles: 3, AddressingMode: 3, PageSensitive: false, Effect: 3},
		&Definition{OpCode: 0x4d, Mnemonic: "EOR", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x4e, Mnemonic: "LSR", Bytes: 3, Cycles: 6, AddressingMode: 3, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x50, Mnemonic: "BVC", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0x51, Mnemonic: "EOR", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		nil,
		nil,
		nil,
		&Definition{OpCode: 0x55, Mnemonic: "EOR", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x56, Mnemonic: "LSR", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x58, Mnemonic: "CLI", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x59, Mnemonic: "EOR", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0x5c, Mnemonic: "skw", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x5d, Mnemonic: "EOR", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x5e, Mnemonic: "LSR", Bytes: 3, Cycles: 7, AddressingMode: 8, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x60, Mnemonic: "RTS", Bytes: 1, Cycles: 6, AddressingMode: 0, PageSensitive: false, Effect: 4},
		&Definition{OpCode: 0x61, Mnemonic: "ADC", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 0},
		nil,
		nil,
		nil,
		&Definition{OpCode: 0x65, Mnemonic: "ADC", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x66, Mnemonic: "ROR", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x68, Mnemonic: "PLA", Bytes: 1, Cycles: 4, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x69, Mnemonic: "ADC", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x6a, Mnemonic: "ROR", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x6b, Mnemonic: "arr", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x6c, Mnemonic: "JMP", Bytes: 3, Cycles: 5, AddressingMode: 5, PageSensitive: false, Effect: 3},
		&Definition{OpCode: 0x6d, Mnemonic: "ADC", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x6e, Mnemonic: "ROR", Bytes: 3, Cycles: 6, AddressingMode: 3, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x70, Mnemonic: "BVS", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0x71, Mnemonic: "ADC", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		nil,
		nil,
		nil,
		&Definition{OpCode: 0x75, Mnemonic: "ADC", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x76, Mnemonic: "ROR", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x78, Mnemonic: "SEI", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x79, Mnemonic: "ADC", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0x7c, Mnemonic: "skw", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x7d, Mnemonic: "ADC", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0x7e, Mnemonic: "ROR", Bytes: 3, Cycles: 7, AddressingMode: 8, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0x80, Mnemonic: "nop", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x81, Mnemonic: "STA", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x82, Mnemonic: "nop", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x83, Mnemonic: "sax", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x84, Mnemonic: "STY", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x85, Mnemonic: "STA", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x86, Mnemonic: "STX", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x87, Mnemonic: "sax", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x88, Mnemonic: "DEY", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0x8a, Mnemonic: "TXA", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x8b, Mnemonic: "xaa", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x8c, Mnemonic: "STY", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x8d, Mnemonic: "STA", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x8e, Mnemonic: "STX", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x8f, Mnemonic: "sax", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x90, Mnemonic: "BCC", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0x91, Mnemonic: "STA", Bytes: 2, Cycles: 6, AddressingMode: 7, PageSensitive: false, Effect: 1},
		nil,
		nil,
		&Definition{OpCode: 0x94, Mnemonic: "STY", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x95, Mnemonic: "STA", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x96, Mnemonic: "STX", Bytes: 2, Cycles: 4, AddressingMode: 11, PageSensitive: false, Effect: 1},
		nil,
		&Definition{OpCode: 0x98, Mnemonic: "TYA", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0x99, Mnemonic: "STA", Bytes: 3, Cycles: 5, AddressingMode: 9, PageSensitive: false, Effect: 1},
		&Definition{OpCode: 0x9a, Mnemonic: "TXS", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0x9d, Mnemonic: "STA", Bytes: 3, Cycles: 5, AddressingMode: 8, PageSensitive: false, Effect: 1},
		nil,
		nil,
		&Definition{OpCode: 0xa0, Mnemonic: "LDY", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa1, Mnemonic: "LDA", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa2, Mnemonic: "LDX", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0xa4, Mnemonic: "LDY", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa5, Mnemonic: "LDA", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa6, Mnemonic: "LDX", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa7, Mnemonic: "lax", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa8, Mnemonic: "TAY", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xa9, Mnemonic: "LDA", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xaa, Mnemonic: "TAX", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0xac, Mnemonic: "LDY", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xad, Mnemonic: "LDA", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xae, Mnemonic: "LDX", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0xb0, Mnemonic: "BCS", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0xb1, Mnemonic: "LDA", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		nil,
		&Definition{OpCode: 0xb3, Mnemonic: "lax", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xb4, Mnemonic: "LDY", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xb5, Mnemonic: "LDA", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xb6, Mnemonic: "LDX", Bytes: 2, Cycles: 4, AddressingMode: 11, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xb7, Mnemonic: "lax", Bytes: 2, Cycles: 4, AddressingMode: 11, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xb8, Mnemonic: "CLV", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xb9, Mnemonic: "LDA", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xba, Mnemonic: "TSX", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0xbc, Mnemonic: "LDY", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xbd, Mnemonic: "LDA", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xbe, Mnemonic: "LDX", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xbf, Mnemonic: "lax", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xc0, Mnemonic: "CPY", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xc1, Mnemonic: "CMP", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0xc4, Mnemonic: "CPY", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xc5, Mnemonic: "CMP", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xc6, Mnemonic: "DEC", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0xc7, Mnemonic: "dcp", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0xc8, Mnemonic: "INY", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xc9, Mnemonic: "CMP", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xca, Mnemonic: "DEX", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xcb, Mnemonic: "axs", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xcc, Mnemonic: "CPY", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xcd, Mnemonic: "CMP", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xce, Mnemonic: "DEC", Bytes: 3, Cycles: 6, AddressingMode: 3, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0xd0, Mnemonic: "BNE", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0xd1, Mnemonic: "CMP", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		nil,
		nil,
		nil,
		&Definition{OpCode: 0xd5, Mnemonic: "CMP", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xd6, Mnemonic: "DEC", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0xd7, Mnemonic: "dcp", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0xd8, Mnemonic: "CLD", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xd9, Mnemonic: "CMP", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0xdc, Mnemonic: "skw", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xdd, Mnemonic: "CMP", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xde, Mnemonic: "DEC", Bytes: 3, Cycles: 7, AddressingMode: 8, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0xe0, Mnemonic: "CPX", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xe1, Mnemonic: "SBC", Bytes: 2, Cycles: 6, AddressingMode: 6, PageSensitive: false, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0xe4, Mnemonic: "CPX", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xe5, Mnemonic: "SBC", Bytes: 2, Cycles: 3, AddressingMode: 4, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xe6, Mnemonic: "INC", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0xe7, Mnemonic: "isc", Bytes: 2, Cycles: 5, AddressingMode: 4, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0xe8, Mnemonic: "INX", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xe9, Mnemonic: "SBC", Bytes: 2, Cycles: 2, AddressingMode: 1, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xea, Mnemonic: "NOP", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		nil,
		&Definition{OpCode: 0xec, Mnemonic: "CPX", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xed, Mnemonic: "SBC", Bytes: 3, Cycles: 4, AddressingMode: 3, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xee, Mnemonic: "INC", Bytes: 3, Cycles: 6, AddressingMode: 3, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0xf0, Mnemonic: "BEQ", Bytes: 2, Cycles: 2, AddressingMode: 2, PageSensitive: true, Effect: 3},
		&Definition{OpCode: 0xf1, Mnemonic: "SBC", Bytes: 2, Cycles: 5, AddressingMode: 7, PageSensitive: true, Effect: 0},
		nil,
		nil,
		nil,
		&Definition{OpCode: 0xf5, Mnemonic: "SBC", Bytes: 2, Cycles: 4, AddressingMode: 10, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xf6, Mnemonic: "INC", Bytes: 2, Cycles: 6, AddressingMode: 10, PageSensitive: false, Effect: 2},
		nil,
		&Definition{OpCode: 0xf8, Mnemonic: "SED", Bytes: 1, Cycles: 2, AddressingMode: 0, PageSensitive: false, Effect: 0},
		&Definition{OpCode: 0xf9, Mnemonic: "SBC", Bytes: 3, Cycles: 4, AddressingMode: 9, PageSensitive: true, Effect: 0},
		nil,
		nil,
		&Definition{OpCode: 0xfc, Mnemonic: "skw", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xfd, Mnemonic: "SBC", Bytes: 3, Cycles: 4, AddressingMode: 8, PageSensitive: true, Effect: 0},
		&Definition{OpCode: 0xfe, Mnemonic: "INC", Bytes: 3, Cycles: 7, AddressingMode: 8, PageSensitive: false, Effect: 2},
		&Definition{OpCode: 0xff, Mnemonic: "isc", Bytes: 3, Cycles: 7, AddressingMode: 8, PageSensitive: false, Effect: 2}}, nil
}
