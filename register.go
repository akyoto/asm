package asm

const registerSizeUnknown = -1

type register struct {
	BaseCodeOffset byte
	BitSize        int
	MustHaveREX    bool
}

var registers = map[string]*register{
	// 8-bit general registers (low)
	"al": &register{0, 8, false},
	"cl": &register{1, 8, false},
	"dl": &register{2, 8, false},
	"bl": &register{3, 8, false},

	// 8-bit general registers (high, without REX prefix)
	"ah": &register{4, 8, false},
	"ch": &register{5, 8, false},
	"dh": &register{6, 8, false},
	"bh": &register{7, 8, false},

	// 8-bit general registers (high, with REX prefix)
	"spl": &register{4, 8, true},
	"bpl": &register{5, 8, true},
	"sil": &register{6, 8, true},
	"dil": &register{7, 8, true},

	// 8-bit general registers
	"r8b":  &register{8, 8, false},
	"r9b":  &register{9, 8, false},
	"r10b": &register{10, 8, false},
	"r11b": &register{11, 8, false},
	"r12b": &register{12, 8, false},
	"r13b": &register{13, 8, false},
	"r14b": &register{14, 8, false},
	"r15b": &register{15, 8, false},

	// 16-bit general registers
	"ax":   &register{0, 16, false},
	"cx":   &register{1, 16, false},
	"dx":   &register{2, 16, false},
	"bx":   &register{3, 16, false},
	"sp":   &register{4, 16, false},
	"bp":   &register{5, 16, false},
	"si":   &register{6, 16, false},
	"di":   &register{7, 16, false},
	"r8w":  &register{8, 16, false},
	"r9w":  &register{9, 16, false},
	"r10w": &register{10, 16, false},
	"r11w": &register{11, 16, false},
	"r12w": &register{12, 16, false},
	"r13w": &register{13, 16, false},
	"r14w": &register{14, 16, false},
	"r15w": &register{15, 16, false},

	// 32-bit general registers
	"eax":  &register{0, 32, false},
	"ecx":  &register{1, 32, false},
	"edx":  &register{2, 32, false},
	"ebx":  &register{3, 32, false},
	"esp":  &register{4, 32, false},
	"ebp":  &register{5, 32, false},
	"esi":  &register{6, 32, false},
	"edi":  &register{7, 32, false},
	"r8d":  &register{8, 32, false},
	"r9d":  &register{9, 32, false},
	"r10d": &register{10, 32, false},
	"r11d": &register{11, 32, false},
	"r12d": &register{12, 32, false},
	"r13d": &register{13, 32, false},
	"r14d": &register{14, 32, false},
	"r15d": &register{15, 32, false},

	// 64-bit general registers
	"rax": &register{0, 64, false},
	"rcx": &register{1, 64, false},
	"rdx": &register{2, 64, false},
	"rbx": &register{3, 64, false},
	"rsp": &register{4, 64, false},
	"rbp": &register{5, 64, false},
	"rsi": &register{6, 64, false},
	"rdi": &register{7, 64, false},
	"r8":  &register{8, 64, false},
	"r9":  &register{9, 64, false},
	"r10": &register{10, 64, false},
	"r11": &register{11, 64, false},
	"r12": &register{12, 64, false},
	"r13": &register{13, 64, false},
	"r14": &register{14, 64, false},
	"r15": &register{15, 64, false},

	// Segment registers
	"es": &register{0, registerSizeUnknown, false},
	"cs": &register{1, registerSizeUnknown, false},
	"ss": &register{2, registerSizeUnknown, false},
	"ds": &register{3, registerSizeUnknown, false},
	"fs": &register{4, registerSizeUnknown, false},
	"gs": &register{5, registerSizeUnknown, false},

	// Floating-point registers
	"st0": &register{0, registerSizeUnknown, false},
	"st1": &register{1, registerSizeUnknown, false},
	"st2": &register{2, registerSizeUnknown, false},
	"st3": &register{3, registerSizeUnknown, false},
	"st4": &register{4, registerSizeUnknown, false},
	"st5": &register{5, registerSizeUnknown, false},
	"st6": &register{6, registerSizeUnknown, false},
	"st7": &register{7, registerSizeUnknown, false},

	// 64-bit mmx registers
	"mm0": &register{0, registerSizeUnknown, false},
	"mm1": &register{1, registerSizeUnknown, false},
	"mm2": &register{2, registerSizeUnknown, false},
	"mm3": &register{3, registerSizeUnknown, false},
	"mm4": &register{4, registerSizeUnknown, false},
	"mm5": &register{5, registerSizeUnknown, false},
	"mm6": &register{6, registerSizeUnknown, false},
	"mm7": &register{7, registerSizeUnknown, false},

	// Control registers
	"cr0": &register{0, registerSizeUnknown, false},
	"cr2": &register{2, registerSizeUnknown, false},
	"cr3": &register{3, registerSizeUnknown, false},
	"cr4": &register{4, registerSizeUnknown, false},

	// Debug registers
	"dr0": &register{0, registerSizeUnknown, false},
	"dr1": &register{1, registerSizeUnknown, false},
	"dr2": &register{2, registerSizeUnknown, false},
	"dr3": &register{3, registerSizeUnknown, false},
	"dr6": &register{6, registerSizeUnknown, false},
	"dr7": &register{7, registerSizeUnknown, false},

	// Test registers
	"tr3": &register{3, registerSizeUnknown, false},
	"tr4": &register{4, registerSizeUnknown, false},
	"tr5": &register{5, registerSizeUnknown, false},
	"tr6": &register{6, registerSizeUnknown, false},
	"tr7": &register{7, registerSizeUnknown, false},
}
