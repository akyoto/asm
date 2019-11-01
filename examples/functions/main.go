package main

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

func main() {
	// Specify program code
	a := asm.New()
	a.MoveRegisterNumber("rdi", 5)
	a.MoveRegisterNumber("rsi", 7)
	a.Call("sum")
	a.Exit(0)

	a.AddLabel("sum")
	a.Println("Hello World")
	a.Return()

	// Compile and save to file
	err := elf.New(a).WriteToFile("program")

	if err != nil {
		log.Fatal(err)
	}
}
