package sections

type Address = uint32

// Strings does string interning and stores the position for each inserted string.
type Strings struct {
	addresses map[string]Address
	raw       []byte
}

// NewStrings creates a new string section.
func NewStrings() *Strings {
	return &Strings{
		addresses: map[string]Address{},
	}
}

// Add adds a string to the section.
func (section *Strings) Add(text string) Address {
	position, exists := section.addresses[text]

	if exists {
		return position
	}

	position = Address(len(section.raw))
	section.addresses[text] = position
	section.raw = append(section.raw, text...)
	section.raw = append(section.raw, 0)
	return position
}

// Bytes returns the entire buffer including all strings
// in the order they were added.
func (section *Strings) Bytes() []byte {
	return section.raw
}

// Count returns the number of added strings.
func (section *Strings) Count() int {
	return len(section.addresses)
}
