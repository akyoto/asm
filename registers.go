package asm

var registers = map[string]*register{
	// 8-bit general registers
	"al": &register{0},
	"cl": &register{1},
	"dl": &register{2},
	"bl": &register{3},
	"ah": &register{4},
	"ch": &register{5},
	"dh": &register{6},
	"bh": &register{7},

	// 16-bit general registers
	"ax": &register{0},
	"cx": &register{1},
	"dx": &register{2},
	"bx": &register{3},
	"sp": &register{4},
	"bp": &register{5},
	"si": &register{6},
	"di": &register{7},

	// 32-bit general registers
	"eax": &register{0},
	"ecx": &register{1},
	"edx": &register{2},
	"ebx": &register{3},
	"esp": &register{4},
	"ebp": &register{5},
	"esi": &register{6},
	"edi": &register{7},

	// 64-bit general registers
	"rax": &register{0},
	"rcx": &register{1},
	"rdx": &register{2},
	"rbx": &register{3},
	"rsp": &register{4},
	"rbp": &register{5},
	"rsi": &register{6},
	"rdi": &register{7},
	"r8":  &register{8},
	"r9":  &register{9},
	"r10": &register{10},
	"r11": &register{11},
	"r12": &register{12},
	"r13": &register{13},
	"r14": &register{14},
	"r15": &register{15},

	// Segment registers
	"es": &register{0},
	"cs": &register{1},
	"ss": &register{2},
	"ds": &register{3},
	"fs": &register{4},
	"gs": &register{5},

	// Floating-point registers
	"st0": &register{0},
	"st1": &register{1},
	"st2": &register{2},
	"st3": &register{3},
	"st4": &register{4},
	"st5": &register{5},
	"st6": &register{6},
	"st7": &register{7},

	// 64-bit mmx registers
	"mm0": &register{0},
	"mm1": &register{1},
	"mm2": &register{2},
	"mm3": &register{3},
	"mm4": &register{4},
	"mm5": &register{5},
	"mm6": &register{6},
	"mm7": &register{7},

	// Control registers
	"cr0": &register{0},
	"cr2": &register{2},
	"cr3": &register{3},
	"cr4": &register{4},

	// Debug registers
	"dr0": &register{0},
	"dr1": &register{1},
	"dr2": &register{2},
	"dr3": &register{3},
	"dr6": &register{6},
	"dr7": &register{7},

	// Test registers
	"tr3": &register{3},
	"tr4": &register{4},
	"tr5": &register{5},
	"tr6": &register{6},
	"tr7": &register{7},
}
