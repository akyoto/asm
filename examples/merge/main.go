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
	a.Call("nicemerge")
	a.Call("exit")
	a.AddLabel("exit")
	a.Exit(0)

	b := asm.New()
	b.AddLabel("hello")
	b.Println("Hello World")
	b.Return()

	c := asm.New()
	c.AddLabel("nicemerge")
	c.Println("Nice merge, isn't it?")
	c.Return()

	// Merge function codes
	a.Merge(b)
	a.Merge(c)

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
