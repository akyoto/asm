package asm

var registerIDs = map[string]byte{
	// 8-bit general registers
	"al": 0, "cl": 1, "dl": 2, "bl": 3, "ah": 4, "ch": 5, "dh": 6, "bh": 7,

	// 16-bit general registers
	"ax": 0, "cx": 1, "dx": 2, "bx": 3, "sp": 4, "bp": 5, "si": 6, "di": 7,

	// 32-bit general registers
	"eax": 0, "ecx": 1, "edx": 2, "ebx": 3, "esp": 4, "ebp": 5, "esi": 6, "edi": 7,

	// 64-bit general registers
	"rax": 0, "rcx": 1, "rdx": 2, "rbx": 3, "rsp": 4, "rbp": 5, "rsi": 6, "rdi": 7,
	"r8": 8, "r9": 9, "r10": 10, "r11": 11, "r12": 12, "r13": 13, "r14": 14, "r15": 15,

	// Segment registers
	"es": 0, "cs": 1, "ss": 2, "ds": 3, "fs": 4, "gs": 5,

	// Floating-point registers
	"st0": 0, "st1": 1, "st2": 2, "st3": 3, "st4": 4, "st5": 5, "st6": 6, "st7": 7,

	// 64-bit mmx registers
	"mm0": 0, "mm1": 1, "mm2": 2, "mm3": 3, "mm4": 4, "mm5": 5, "mm6": 6, "mm7": 7,

	// Control registers
	"cr0": 0, "cr2": 2, "cr3": 3, "cr4": 4,

	// Debug registers
	"dr0": 0, "dr1": 1, "dr2": 2, "dr3": 3, "dr6": 6, "dr7": 7,

	// Test registers
	"tr3": 3, "tr4": 4, "tr5": 5, "tr6": 6, "tr7": 7,
}
