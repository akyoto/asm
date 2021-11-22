# {name}

{go:header}

An x86-64 assembler written in Go. It is used by the [Q programming language](https://github.com/akyoto/q) for machine code generation.

It was born out of the need for a highly performant assembler that had next to no overhead when compiling instructions to bytecode. When I started this project, I had no idea of how x86 instructions were constructed and I was forced to do a lot of reverse engineering via NASM.

Hopefully this repository helps somebody who is learning about x86 assembly and its bytecode format.

There are a few [examples](https://github.com/akyoto/asm/tree/master/examples) to get an idea of the API I was aiming for.

## x86 bytecode

An x86 program consists of a list of instructions. All of these instructions are built with the following format:

| Name                | Size in bytes | Required? |
|---------------------|---------------|-----------|
| Instruction prefix  | 1             |           |
| Address size prefix | 1             |           |
| Operand size prefix | 1             |           |
| Segment override    | 1             |           |
| OP code             | 1-2           | required  |
| Mod/RM              | 1             |           |
| SIB                 | 1             |           |
| Displacement        | 1-4           |           |
| Immediate           | 1-4           |           |

Out of these only the actual OP code which decides the instruction to execute is required. The remaining components depend on what instruction and what kind of parameters you have.

The Mod/RM byte has 3 sub components:

| Name | Size in bits |
|------|--------------|
| Mod  | 2            |
| Reg  | 3            |
| R/M  | 3            |

You can build it like this:

```go
(mod << 6) | (reg << 3) | rm
```

The SIB byte has the same format:

| Name  | Size in bits |
|-------|--------------|
| Scale | 2            |
| Index | 3            |
| Base  | 3            |

You can build it like this:

```go
(scale << 6) | (index << 3) | base
```

The `opcode` directory has a few helper functions to construct these components.

## Registers

|  8B  |  4B  |  2B  |  1B  |
|------|------|------|------|
| rax  | eax  | ax   | al   |
| rcx  | ecx  | cx   | cl   |
| rdx  | edx  | dx   | dl   |
| rbx  | ebx  | bx   | bl   |
| rsi  | esi  | si   | sil  |
| rdi  | edi  | di   | dil  |
| rsp  | esp  | sp   | spl  |
| rbp  | ebp  | bp   | bpl  |
| r8   | r8d  | r8w  | r8b  |
| r9   | r9d  | r9w  | r9b  |
| r10  | r10d | r10w | r10b |
| r11  | r11d | r11w | r11b |
| r12  | r12d | r12w | r12b |
| r13  | r13d | r13w | r13b |
| r14  | r14d | r14w | r14b |
| r15  | r15d | r15w | r15b |

## Status

This project is currently work in progress. Contributions are welcome.

{go:footer}
