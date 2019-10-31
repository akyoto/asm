package sections

// Pointer stores a relative memory address reference (relative to the section)
// that we can later turn into an absolute one.
// Address:  The offset inside the section.
// Position: The machine code offset where the address was inserted.
type Pointer struct {
	Address  int64
	Position int
}
