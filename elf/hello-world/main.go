package main

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

func main() {
	a := asm.New()
	a.Print("Hello World\n")
	a.Print("Nice day, isn't it?\n")
	a.Exit(0)

	program := elf.New(a.Bytes(), a.StringTable, a.SectionPointers)
	err := program.WriteToFile("hello")

	if err != nil {
		log.Fatal(err)
	}
}
