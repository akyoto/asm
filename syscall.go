package asm

import "github.com/akyoto/asm/syscall"

// Syscall calls a kernel function with the given parameters.
func (a *Assembler) Syscall(parameters ...interface{}) {
	for count, parameter := range parameters {
		a.MoveRegisterNumber(syscall.Registers[count], parameter)
	}

	a.WriteBytes(0x0f, 0x05)
}

// Print prints a message on the terminal.
func (a *Assembler) Print(msg string) {
	a.Syscall(syscall.Write, int32(1), msg, int32(len(msg)))
}

// Print prints a message followed by a new line on the terminal.
func (a *Assembler) Println(msg string) {
	a.Print(msg + "\n")
}

// Open opens a file.
func (a *Assembler) Open(fileName string) {
	a.Syscall(syscall.Open, 2, fileName, int32(0102), int32(0666))
}

// Exit terminates the program.
func (a *Assembler) Exit(code int32) {
	a.Syscall(syscall.Exit, code)
}
