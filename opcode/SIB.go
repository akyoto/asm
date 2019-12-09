package opcode

// SIB is used to generate an SIB byte.
// - scale: 2 bits
// - index: 3 bits
// - base:  3 bits
func SIB(scale byte, index byte, base byte) byte {
	return (scale << 6) | (index << 3) | base
}
