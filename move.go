package asm

import (
	"encoding/binary"
	"log"
	"math"

	"github.com/akyoto/asm/opcode"
)

// MoveRegisterNumber moves a number into the given register.
func (a *Assembler) MoveRegisterNumber(registerNameTo string, number interface{}) {
	baseCode := byte(0xb8)
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	if registerTo.BitSize == 8 {
		baseCode = 0xb0
	}

	if registerTo.BitSize == 16 {
		a.WriteBytes(0x66)
	}

	operandBitSize := 0
	numberConverted := uint64(0)
	isString := false

	switch value := number.(type) {
	case string:
		operandBitSize = 32
		numberConverted = 0
		isString = true

	case int64:
		operandBitSize = 64
		numberConverted = uint64(value)

	case int:
		operandBitSize = 64
		numberConverted = uint64(value)

	case int32:
		operandBitSize = 32
		numberConverted = uint64(value)

	case int16:
		operandBitSize = 16
		numberConverted = uint64(value)

	case byte:
		operandBitSize = 8
		numberConverted = uint64(value)

	default:
		log.Fatalf("Unsupported type: %v", value)
	}

	registerBitSize := registerTo.BitSize

	if a.EnableOptimizer && registerBitSize == 64 && numberConverted < math.MaxUint32 {
		registerBitSize = 32
		operandBitSize = 32
	}

	if registerBitSize != operandBitSize {
		log.Printf("Register size (%d) doesn't match operand size (%d)", registerBitSize, operandBitSize)
	}

	bitSize := registerBitSize

	// 64-bit register
	w := byte(0)

	if bitSize == 64 {
		w = 1
	}

	// Are we accessing any of the 64-bit only registers (r8 up to r15)?
	b := byte(0)

	if registerTo.BaseCodeOffset >= 8 {
		b = 1
	}

	// REX
	if w != 0 || b != 0 || registerTo.MustHaveREX {
		a.WriteBytes(opcode.REX(w, 0, 0, b))
	}

	// Base code
	a.WriteBytes(baseCode + registerTo.BaseCodeOffset%8)

	// Add string address after the instruction code
	if isString {
		numberConverted = uint64(a.AddString(number.(string)))
	}

	// Number
	var buffer []byte

	switch bitSize {
	case 64:
		buffer = make([]byte, 8)
		binary.LittleEndian.PutUint64(buffer, numberConverted)

	case 32:
		buffer = make([]byte, 4)
		binary.LittleEndian.PutUint32(buffer, uint32(numberConverted))

	case 16:
		buffer = make([]byte, 2)
		binary.LittleEndian.PutUint16(buffer, uint16(numberConverted))

	case 8:
		buffer = []byte{byte(numberConverted)}
	}

	_, _ = a.Write(buffer)
}

// MoveRegisterRegister moves a register value into another register.
func (a *Assembler) MoveRegisterRegister(registerNameTo string, registerNameFrom string) {
	baseCode := byte(0x89)
	registerTo, exists := registers[registerNameTo]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameTo)
	}

	registerFrom, exists := registers[registerNameFrom]

	if !exists {
		log.Fatal("Unknown register name: " + registerNameFrom)
	}

	if registerTo.BitSize == 64 {
		r := byte(0)
		b := byte(0)

		if registerFrom.BaseCodeOffset >= 8 {
			r = 1
		}

		if registerTo.BaseCodeOffset >= 8 {
			b = 1
		}

		a.WriteBytes(opcode.REX(1, r, 0, b))
	}

	a.WriteBytes(baseCode)
	a.WriteBytes(opcode.ModRM(0b11, registerFrom.BaseCodeOffset, registerTo.BaseCodeOffset))
}
