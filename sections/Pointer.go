package sections

// Pointer64 stores a relative memory address reference (relative to the section)
// that we can later turn into an absolute one.
// Address:  The offset inside the section.
// Position: The machine code offset where the address was inserted.
type Pointer64 struct {
	Address  int64
	Position int32
}

// Pointer32 is the same as Pointer64 except it uses 32-bit addresses.
type Pointer32 struct {
	Address  int32
	Position int32
}
