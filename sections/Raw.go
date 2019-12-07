package sections

import "bytes"

// Raw stores the position for each inserted byte slice.
type Raw struct {
	raw bytes.Buffer
}

// NewRaw creates a new raw data section.
func NewRaw() *Raw {
	return &Raw{}
}

// Add adds a byte slice to the section.
func (section *Raw) Add(data []byte) Address {
	position := Address(section.raw.Len())
	section.raw.Write(data)
	return position
}

// Bytes returns the entire buffer including all data
// in the order they were added.
func (section *Raw) Bytes() []byte {
	return section.raw.Bytes()
}
