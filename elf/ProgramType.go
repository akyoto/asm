package elf

type ProgramType int32

const (
	ProgramTypeNULL    ProgramType = 0
	ProgramTypeLOAD    ProgramType = 1
	ProgramTypeDYNAMIC ProgramType = 2
	ProgramTypeINTERP  ProgramType = 3
	ProgramTypeNOTE    ProgramType = 4
	ProgramTypeSHLIB   ProgramType = 5
	ProgramTypePHDR    ProgramType = 6
	ProgramTypeTLS     ProgramType = 7 // Thread local storage segment
)
