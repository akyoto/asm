package asm

const registerSizeUnknown = -1

type register struct {
	BaseCodeOffset byte
	BitSize        int
	MustHaveREX    bool
}

var registers = map[string]*register{
	// 8-bit general registers (low)
	"al": {0, 8, false},
	"cl": {1, 8, false},
	"dl": {2, 8, false},
	"bl": {3, 8, false},

	// 8-bit general registers (high, without REX prefix)
	"ah": {4, 8, false},
	"ch": {5, 8, false},
	"dh": {6, 8, false},
	"bh": {7, 8, false},

	// 8-bit general registers (high, with REX prefix)
	"spl": {4, 8, true},
	"bpl": {5, 8, true},
	"sil": {6, 8, true},
	"dil": {7, 8, true},

	// 8-bit general registers
	"r8b":  {8, 8, false},
	"r9b":  {9, 8, false},
	"r10b": {10, 8, false},
	"r11b": {11, 8, false},
	"r12b": {12, 8, false},
	"r13b": {13, 8, false},
	"r14b": {14, 8, false},
	"r15b": {15, 8, false},

	// 16-bit general registers
	"ax":   {0, 16, false},
	"cx":   {1, 16, false},
	"dx":   {2, 16, false},
	"bx":   {3, 16, false},
	"sp":   {4, 16, false},
	"bp":   {5, 16, false},
	"si":   {6, 16, false},
	"di":   {7, 16, false},
	"r8w":  {8, 16, false},
	"r9w":  {9, 16, false},
	"r10w": {10, 16, false},
	"r11w": {11, 16, false},
	"r12w": {12, 16, false},
	"r13w": {13, 16, false},
	"r14w": {14, 16, false},
	"r15w": {15, 16, false},

	// 32-bit general registers
	"eax":  {0, 32, false},
	"ecx":  {1, 32, false},
	"edx":  {2, 32, false},
	"ebx":  {3, 32, false},
	"esp":  {4, 32, false},
	"ebp":  {5, 32, false},
	"esi":  {6, 32, false},
	"edi":  {7, 32, false},
	"r8d":  {8, 32, false},
	"r9d":  {9, 32, false},
	"r10d": {10, 32, false},
	"r11d": {11, 32, false},
	"r12d": {12, 32, false},
	"r13d": {13, 32, false},
	"r14d": {14, 32, false},
	"r15d": {15, 32, false},

	// 64-bit general registers
	"rax": {0, 64, false},
	"rcx": {1, 64, false},
	"rdx": {2, 64, false},
	"rbx": {3, 64, false},
	"rsp": {4, 64, false},
	"rbp": {5, 64, false},
	"rsi": {6, 64, false},
	"rdi": {7, 64, false},
	"r8":  {8, 64, false},
	"r9":  {9, 64, false},
	"r10": {10, 64, false},
	"r11": {11, 64, false},
	"r12": {12, 64, false},
	"r13": {13, 64, false},
	"r14": {14, 64, false},
	"r15": {15, 64, false},

	// Segment registers
	"es": {0, registerSizeUnknown, false},
	"cs": {1, registerSizeUnknown, false},
	"ss": {2, registerSizeUnknown, false},
	"ds": {3, registerSizeUnknown, false},
	"fs": {4, registerSizeUnknown, false},
	"gs": {5, registerSizeUnknown, false},

	// Floating-point registers
	"st0": {0, registerSizeUnknown, false},
	"st1": {1, registerSizeUnknown, false},
	"st2": {2, registerSizeUnknown, false},
	"st3": {3, registerSizeUnknown, false},
	"st4": {4, registerSizeUnknown, false},
	"st5": {5, registerSizeUnknown, false},
	"st6": {6, registerSizeUnknown, false},
	"st7": {7, registerSizeUnknown, false},

	// 64-bit mmx registers
	"mm0": {0, registerSizeUnknown, false},
	"mm1": {1, registerSizeUnknown, false},
	"mm2": {2, registerSizeUnknown, false},
	"mm3": {3, registerSizeUnknown, false},
	"mm4": {4, registerSizeUnknown, false},
	"mm5": {5, registerSizeUnknown, false},
	"mm6": {6, registerSizeUnknown, false},
	"mm7": {7, registerSizeUnknown, false},

	// Control registers
	"cr0": {0, registerSizeUnknown, false},
	"cr2": {2, registerSizeUnknown, false},
	"cr3": {3, registerSizeUnknown, false},
	"cr4": {4, registerSizeUnknown, false},

	// Debug registers
	"dr0": {0, registerSizeUnknown, false},
	"dr1": {1, registerSizeUnknown, false},
	"dr2": {2, registerSizeUnknown, false},
	"dr3": {3, registerSizeUnknown, false},
	"dr6": {6, registerSizeUnknown, false},
	"dr7": {7, registerSizeUnknown, false},

	// Test registers
	"tr3": {3, registerSizeUnknown, false},
	"tr4": {4, registerSizeUnknown, false},
	"tr5": {5, registerSizeUnknown, false},
	"tr6": {6, registerSizeUnknown, false},
	"tr7": {7, registerSizeUnknown, false},
}
