package asm

const registerSizeUnknown = -1

type register struct {
	BaseCodeOffset byte
	BitSize        int
}

var registers = map[string]*register{
	// 8-bit general registers
	"al": &register{0, 8},
	"cl": &register{1, 8},
	"dl": &register{2, 8},
	"bl": &register{3, 8},
	"ah": &register{4, 8},
	"ch": &register{5, 8},
	"dh": &register{6, 8},
	"bh": &register{7, 8},

	// 16-bit general registers
	"ax": &register{0, 16},
	"cx": &register{1, 16},
	"dx": &register{2, 16},
	"bx": &register{3, 16},
	"sp": &register{4, 16},
	"bp": &register{5, 16},
	"si": &register{6, 16},
	"di": &register{7, 16},

	// 16-bit general registers in x64
	"r8w":  &register{8, 16},
	"r9w":  &register{9, 16},
	"r10w": &register{10, 16},
	"r11w": &register{11, 16},
	"r12w": &register{12, 16},
	"r13w": &register{13, 16},
	"r14w": &register{14, 16},
	"r15w": &register{15, 16},

	// 32-bit general registers
	"eax": &register{0, 32},
	"ecx": &register{1, 32},
	"edx": &register{2, 32},
	"ebx": &register{3, 32},
	"esp": &register{4, 32},
	"ebp": &register{5, 32},
	"esi": &register{6, 32},
	"edi": &register{7, 32},

	// 64-bit general registers
	"rax": &register{0, 64},
	"rcx": &register{1, 64},
	"rdx": &register{2, 64},
	"rbx": &register{3, 64},
	"rsp": &register{4, 64},
	"rbp": &register{5, 64},
	"rsi": &register{6, 64},
	"rdi": &register{7, 64},
	"r8":  &register{8, 64},
	"r9":  &register{9, 64},
	"r10": &register{10, 64},
	"r11": &register{11, 64},
	"r12": &register{12, 64},
	"r13": &register{13, 64},
	"r14": &register{14, 64},
	"r15": &register{15, 64},

	// Segment registers
	"es": &register{0, registerSizeUnknown},
	"cs": &register{1, registerSizeUnknown},
	"ss": &register{2, registerSizeUnknown},
	"ds": &register{3, registerSizeUnknown},
	"fs": &register{4, registerSizeUnknown},
	"gs": &register{5, registerSizeUnknown},

	// Floating-point registers
	"st0": &register{0, registerSizeUnknown},
	"st1": &register{1, registerSizeUnknown},
	"st2": &register{2, registerSizeUnknown},
	"st3": &register{3, registerSizeUnknown},
	"st4": &register{4, registerSizeUnknown},
	"st5": &register{5, registerSizeUnknown},
	"st6": &register{6, registerSizeUnknown},
	"st7": &register{7, registerSizeUnknown},

	// 64-bit mmx registers
	"mm0": &register{0, registerSizeUnknown},
	"mm1": &register{1, registerSizeUnknown},
	"mm2": &register{2, registerSizeUnknown},
	"mm3": &register{3, registerSizeUnknown},
	"mm4": &register{4, registerSizeUnknown},
	"mm5": &register{5, registerSizeUnknown},
	"mm6": &register{6, registerSizeUnknown},
	"mm7": &register{7, registerSizeUnknown},

	// Control registers
	"cr0": &register{0, registerSizeUnknown},
	"cr2": &register{2, registerSizeUnknown},
	"cr3": &register{3, registerSizeUnknown},
	"cr4": &register{4, registerSizeUnknown},

	// Debug registers
	"dr0": &register{0, registerSizeUnknown},
	"dr1": &register{1, registerSizeUnknown},
	"dr2": &register{2, registerSizeUnknown},
	"dr3": &register{3, registerSizeUnknown},
	"dr6": &register{6, registerSizeUnknown},
	"dr7": &register{7, registerSizeUnknown},

	// Test registers
	"tr3": &register{3, registerSizeUnknown},
	"tr4": &register{4, registerSizeUnknown},
	"tr5": &register{5, registerSizeUnknown},
	"tr6": &register{6, registerSizeUnknown},
	"tr7": &register{7, registerSizeUnknown},
}
