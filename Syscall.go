package asm

import (
	"github.com/akyoto/asm/syscall"
)

// Print prints a message on the terminal.
func (a *Assembler) Print(msg string) {
	address := a.AddData([]byte(msg))

	a.MoveRegisterNumber(syscall.Registers[0], uint64(syscall.Write))
	a.MoveRegisterNumber(syscall.Registers[1], 1)
	a.MoveRegisterAddress(syscall.Registers[2], address)
	a.MoveRegisterNumber(syscall.Registers[3], uint64(len(msg)))
	a.Syscall()
}

// Print prints a message followed by a new line on the terminal.
func (a *Assembler) Println(msg string) {
	a.Print(msg + "\n")
}

// Exit terminates the program.
func (a *Assembler) Exit(code int32) {
	a.MoveRegisterNumber(syscall.Registers[0], uint64(syscall.Exit))
	a.MoveRegisterNumber(syscall.Registers[1], uint64(code))
	a.Syscall()
}

// Open opens a file.
func (a *Assembler) Open(fileName string) {
	address := a.AddData([]byte(fileName))

	a.MoveRegisterNumber(syscall.Registers[0], uint64(syscall.Open))
	a.MoveRegisterNumber(syscall.Registers[1], 2)
	a.MoveRegisterAddress(syscall.Registers[2], address)
	a.MoveRegisterNumber(syscall.Registers[3], uint64(0102))
	a.MoveRegisterNumber(syscall.Registers[4], uint64(0666))
	a.Syscall()
}

// Syscall executes a syscall.
func (a *Assembler) Syscall() {
	a.WriteBytes(0x0f, 0x05)
}
