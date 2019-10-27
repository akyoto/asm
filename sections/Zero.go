package sections

// Zero stores the position for each inserted 0-initialized global variable.
type Zero struct {
	length int64
}

// NewZero creates a new zero initialized data section (bss).
func NewZero() *Zero {
	return &Zero{}
}

// Add adds the given number of bytes to the section.
func (section *Zero) Add(count int64) int64 {
	position := int64(section.length)
	section.length += count
	return position
}
