package main

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

func main() {
	// Specify program code
	a := asm.New()
	a.Println("Hello World")
	a.Exit(0)

	// Compile
	err := a.Compile()

	if err != nil {
		log.Fatal(err)
	}

	// Save to file
	err = elf.New(a).WriteToFile("program")

	if err != nil {
		log.Fatal(err)
	}
}
