package asm

// CPUID is used to query CPU relevant data based on the contents of EAX.
func (a *Assembler) CPUID() {
	a.WriteBytes(0x0f, 0xa2)
}

// ReadTimeStampCounterAndProcessorID is used for performance benchmarks.
func (a *Assembler) ReadTimeStampCounterAndProcessorID() {
	a.WriteBytes(0x0f, 0x01, 0xf9)
}

// EndBr64 is a security related instruction enabling CET (Control-flow Enforcement Technology).
// It ensures that indirect branches go to a valid location.
func (a *Assembler) EndBr64() {
	a.WriteBytes(0xf3, 0x0f, 0x1e, 0xfa)
}
