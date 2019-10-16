package asm

import "bytes"

type stringTable struct {
	virtualOffset int64
	addresses     map[string]int64
	raw           bytes.Buffer
}

func newStringTable() *stringTable {
	return &stringTable{
		virtualOffset: 0x400000 + 0xee,
		addresses:     make(map[string]int64),
	}
}

func (table *stringTable) Add(text string) int64 {
	position, exists := table.addresses[text]

	if exists {
		return table.virtualOffset + position
	}

	position = int64(table.raw.Len())
	table.addresses[text] = position
	table.raw.WriteString(text)
	return table.virtualOffset + position
}
