package main

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

func main() {
	a := asm.New()
	a.Println("Hello World")
	a.Println("Nice day, isn't it?")
	a.Exit(0)

	program := elf.New(a.Bytes(), a.StringTable, a.SectionPointers)
	err := program.WriteToFile("hello")

	if err != nil {
		log.Fatal(err)
	}
}
