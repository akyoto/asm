package main

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

func main() {
	a := asm.New()
	a.Println("Hello World")
	a.Exit(0)

	program := elf.New(a.Bytes(), a.Strings, a.StringPointers)
	err := program.WriteToFile("program")

	if err != nil {
		log.Fatal(err)
	}
}
