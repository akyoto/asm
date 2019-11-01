# asm

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

An x86-64 assembler written in Go.

## Architectures

- [x] Linux x86-64 (ELF binaries)
- [ ] ...

## Examples

See [examples](https://github.com/akyoto/asm/tree/master/examples).

## Reference

### Registers

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

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://eduardurbach.com) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

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
