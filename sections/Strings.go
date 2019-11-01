package sections

import "bytes"

// Strings does string interning and stores the position for each inserted string.
type Strings struct {
	addresses map[string]int64
	raw       bytes.Buffer
}

// NewStrings creates a new string section.
func NewStrings() *Strings {
	return &Strings{
		addresses: make(map[string]int64),
	}
}

// Add adds a string to the section.
func (section *Strings) Add(text string) int64 {
	position, exists := section.addresses[text]

	if exists {
		return position
	}

	position = int64(section.raw.Len())
	section.addresses[text] = position
	section.raw.WriteString(text)
	section.raw.WriteByte(0)
	return position
}

// Bytes returns the entire buffer including all strings
// in the order they were added.
func (section *Strings) Bytes() []byte {
	return section.raw.Bytes()
}

// Count returns the number of added strings.
func (section *Strings) Count() int {
	return len(section.addresses)
}
