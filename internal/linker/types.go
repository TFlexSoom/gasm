package linker

type ExternRef string
type LibraryName string
type Path string

type BinaryFileFlag int

const (
	STRIPPED_RELOCATION BinaryFileFlag = iota
	IS_RELOCATABLE
	STRIPPED_LINE_NUMBERS
	STRIPPED_SYMBOLS
	LITTLE_ENDIAN
	BIG_ENDIAN_MARKED
	DUPLICATE_SYMBOLS_REMOVED
)

type BinarySectionFlag int

const (
	BSF_REGULAR BinarySectionFlag = iota
	BSF_DUMMY
	BSF_NO_LOAD
	BSF_GROUP
	BSF_PADDING
	BSF_COPY
	BSF_TEXT
	BSF_DATA
	BSF_BSS
	BSF_BLOCK
	BSF_PASS
	BSF_CLINK
	BSF_VECTOR
	BSF_PADDED
)

type BinarySection interface {
	Name() string
	Size() uint64
	Content() []byte
	Flags() map[BinarySectionFlag]bool
}

type RelocationType uint16

const (
	RT_ADD RelocationType = iota
	RT_SUB
	RT_NEG
	RT_MPY
	RT_DIV
	RT_MOD
	RT_SR
	RT_ASR
	RT_SL
	RT_AND
	RT_OR
	RT_XOR
	RT_NOTB
	RT_ULDFLD
	RT_SLDFLD
	RT_USTFLD
	RT_SSTFLD
	RT_PUSH
	RT_PUSHSK
	RT_PUSHUK
	RT_PUSHPC
	RT_DUP
	RT_XSTFLD
	RT_PUSHSV
	RT_ABS
	RT_RELBYTE
	RT_RELWORD
	RT_REL24
	RT_RELLONG
)

type RelocationEntry interface {
	VirtualAddress() uintptr
	SymTableIndex() uint16
	RelocationType() RelocationType
}

type StorageClass uint8

const (
	SC_NULL uint8 = iota
	SC_AUTO
	SC_EXT
	SC_STAT
	SC_REG
	SC_EXTREF
	SC_LABEL
	SC_ULABEL
	SC_MOS
	SC_ARG
	SC_STRTAG
	SC_MOU
	SC_UNTAG
	SC_TPDEF
	SC_USTATIC
	SC_ENTAG
	SC_MOE
	SC_REGPARAM
	SC_FIELD
	SC_UEXT
	SC_STATLAB
	SC_EXTLAB

	SC_VARARG = 27

	SC_BLOCK = 100
	SC_FCN   = 101
	SC_EOS   = 102
	SC_FILE  = 103
	SC_LINE  = 104
)

type SymbolEntry interface {
	Name() string
	Value() []byte
	SectionNo() int16
	StorageClass() StorageClass
	AuxEntries() bool
	IsAbsolute() bool
	IsDefined() bool
}

// .text -> Address of .text section
// .data -> Address of .data section
// etext -> Next Available after .text
// ...

type BinaryFile interface {
	//-- File Header
	VersionNo() uint16
	NumberOfSections() uint16
	Timestamp() uint32
	// symbol start
	NumOfSymbols() uint32
	Flags() map[BinaryFileFlag]bool
	HasHeader() bool
	TargetMagicNumber() uint16

	//-- Optional File Header
	FileMagicNumber() uint16
	VersionStamp() uint16
	ExecSize() uint64
	DataSize() uint64
	BssSize() uint64

	// entryPoint

	// beginning of executable code
	// beginning of initialized data

	// section header structure
	Sections() []BinarySection
	Relocations() []RelocationEntry
	Symbols() []SymbolEntry
}
