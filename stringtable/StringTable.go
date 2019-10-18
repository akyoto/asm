package stringtable

import "bytes"

// StringTable does string interning and stores the position for each inserted string.
type StringTable struct {
	addresses map[string]int64
	raw       bytes.Buffer
}

// New creates a new string table.
func New() *StringTable {
	return &StringTable{
		addresses: make(map[string]int64),
	}
}

// Add adds a string to the table.
func (table *StringTable) Add(text string) int64 {
	position, exists := table.addresses[text]

	if exists {
		return position
	}

	position = int64(table.raw.Len())
	table.addresses[text] = position
	table.raw.WriteString(text)
	table.raw.WriteByte(0)
	return position
}

// Bytes returns the entire buffer including all strings
// in the order they were added.
func (table *StringTable) Bytes() []byte {
	return table.raw.Bytes()
}
