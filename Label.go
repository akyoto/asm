package asm

// AddLabelAt adds a label for the current instruction address.
func (a *Assembler) AddLabel(name string) {
	a.AddLabelAt(name, a.Position())
}

// AddLabelAt adds a label for the given address.
func (a *Assembler) AddLabelAt(name string, address uint32) {
	a.labels[name] = address
}
