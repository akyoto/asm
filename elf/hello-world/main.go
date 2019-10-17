package main

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

func main() {
	a := asm.New()
	a.Print("Hello World\n")
	a.Exit(0)

	program := elf.New(a.Bytes())
	err := program.WriteToFile("hello")

	if err != nil {
		log.Fatal(err)
	}
}
