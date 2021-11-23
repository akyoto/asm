# asm

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

An x86-64 assembler written in Go. It is used by the [Q programming language](https://github.com/akyoto/q) for machine code generation.

It was born out of the need for a highly performant assembler that had next to no overhead when compiling instructions to bytecode. When I started this project, I had no idea of how x86 instructions were constructed and I was forced to do a lot of reverse engineering via NASM.

Hopefully this repository helps someone who is learning about x86 assembly and its bytecode format.

There are a few [examples](https://github.com/akyoto/asm/tree/master/examples) to get an idea of the API I was aiming for.

## x86-64 bytecode

An x86-64 program consists of a list of instructions. All of these instructions are built with the following format:

| Name            | Size in bytes | Required? |
|-----------------|---------------|-----------|
| Legacy prefixes | 1-4           |           |
| OP code         | 1-4           | required  |
| Mod/RM          | 1             |           |
| SIB             | 1             |           |
| Displacement    | 1-8           |           |
| Immediate       | 1-8           |           |

Out of these only the actual OP code which decides the instruction to execute is required. The remaining components depend on what instruction and what kind of parameters you have.

The maximum size for a single instruction is limited to 15 bytes.

### Mod/RM

The Mod/RM byte has the following format:

| Name | Size in bits |
|------|--------------|
| Mod  | 2            |
| Reg  | 3            |
| RM   | 3            |

```go
(mod << 6) | (reg << 3) | rm
```

### SIB

The SIB byte has the same format as Mod/RM, just with different meanings:

| Name  | Size in bits |
|-------|--------------|
| Scale | 2            |
| Index | 3            |
| Base  | 3            |

```go
(scale << 6) | (index << 3) | base
```

The `opcode` directory has a few helper functions to construct these components.

## Registers

The following is a list of register names you can use. I decided to stick with the original names instead of r0-r7 for rax-rbp. I might still switch to r0-r7 for the future and enable the old names as synonyms.

| 64 bit | 32 bit | 16 bit | 8 bit |
|--------|--------|--------|-------|
| rax    | eax    | ax     | al    |
| rcx    | ecx    | cx     | cl    |
| rdx    | edx    | dx     | dl    |
| rbx    | ebx    | bx     | bl    |
| rsi    | esi    | si     | sil   |
| rdi    | edi    | di     | dil   |
| rsp    | esp    | sp     | spl   |
| rbp    | ebp    | bp     | bpl   |
| r8     | r8d    | r8w    | r8b   |
| r9     | r9d    | r9w    | r9b   |
| r10    | r10d   | r10w   | r10b  |
| r11    | r11d   | r11w   | r11b  |
| r12    | r12d   | r12w   | r12b  |
| r13    | r13d   | r13w   | r13b  |
| r14    | r14d   | r14w   | r14b  |
| r15    | r15d   | r15w   | r15b  |

## Resources

* https://godbolt.org
* https://www.felixcloutier.com/x86
* https://wiki.osdev.org/X86-64_Instruction_Encoding
* https://www.systutorials.com/beginners-guide-x86-64-instruction-encoding/
* https://blog.yossarian.net/2020/06/13/How-x86_64-addresses-memory
* https://cs.brown.edu/courses/cs033/docs/guides/x64_cheatsheet.pdf
* https://www.cs.cmu.edu/~fp/courses/15213-s07/resources.html
* https://www.agner.org/optimize/optimizing_assembly.pdf
* https://www.intel.com/content/www/us/en/developer/articles/technical/intel-sdm.html

[godoc-image]: https://godoc.org/github.com/akyoto/asm?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/asm
[report-image]: https://goreportcard.com/badge/github.com/akyoto/asm
[report-url]: https://goreportcard.com/report/github.com/akyoto/asm
[tests-image]: https://cloud.drone.io/api/badges/akyoto/asm/status.svg
[tests-url]: https://cloud.drone.io/akyoto/asm
[coverage-image]: https://codecov.io/gh/akyoto/asm/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/asm
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
