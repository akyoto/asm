package elf

type SectionFlags int64

const (
	SectionFlagsWritable   SectionFlags = 0x1
	SectionFlagsAllocate   SectionFlags = 0x2
	SectionFlagsExecutable SectionFlags = 0x4
)
