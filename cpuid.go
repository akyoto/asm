package asm

// CPUID is used to query CPU relevant data based on the contents of EAX.
func (a *Assembler) CPUID() {
	a.WriteBytes(0x0f, 0xa2)
}

// ReadTimeStampCounterAndProcessorID is used for performance benchmarks.
func (a *Assembler) ReadTimeStampCounterAndProcessorID() {
	a.WriteBytes(0x0f, 0x01, 0xf9)
}
