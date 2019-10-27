package asm

// ModRM is used to generate a ModRM suffix.
// - mod: 2 bits
// - reg: 3 bits
// - rm:  3 bits
func ModRM(mod byte, reg byte, rm byte) byte {
	return (mod << 6) | (reg << 3) | rm
}
