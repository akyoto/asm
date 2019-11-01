package main

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

func main() {
	// Specify program code
	a := asm.New()

	a.Call("hello")
	a.Call("niceday")
	a.Call("exit")

	a.AddLabel("niceday")
	a.Println("Nice day, isn't it?")
	a.Return()

	a.AddLabel("hello")
	a.Println("Hello World")
	a.Return()

	a.AddLabel("exit")
	a.Exit(0)

	// Compile and save to file
	err := elf.New(a).WriteToFile("program")

	if err != nil {
		log.Fatal(err)
	}
}
