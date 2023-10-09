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
	REGULAR BinarySectionFlag = iota
	DUMMY
	NO_LOAD
	GROUP
	PADDING
	COPY
	TEXT
	DATA
	BSS
	BLOCK
	PASS
	CLINK
	VECTOR
	PADDED
)

type BinarySection interface {
	name() string
	size() uint64
	content() []byte
	relocations() []bool
	flags() map[BinarySectionFlag]bool
}

type RelocationType uint16

const (
	ADD RelocationType = iota
	SUB
	NEG
	MPY
	DIV
	MOD
	SR
	ASR
	SL
	AND
	OR
	XOR
	NOTB
	ULDFLD
	SLDFLD
)

type RelocationEntry interface {
	virtualAddress() uintptr
	symTableIndex() uint16
	relocationType() RelocationType
}

type BinaryFile interface {
	versionNo() uint16
	numberOfSections() uint16
	timestamp() uint32
	// symbol start
	numOfSymbols() uint32
	flags() map[BinaryFileFlag]bool
	hasHeader() bool
	targetMagicNumber() uint16

	// Optional File Header
	fileMagicNumber() uint16
	versionStamp() uint16
	execSize() uint64
	dataSize() uint64
	bssSize() uint64
	// entryPoint
	// beginning of executable code
	// beginning of initialized data

	// section header structure
	sections() []BinarySection
}
