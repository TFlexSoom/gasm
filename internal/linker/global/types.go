package linker

type ExternRef string
type LibraryName string
type Path string

type BinarySection interface {
	Name() string
	Start() uintptr
	Size() uint64
	Content() []byte
	Flags() map[BinarySectionFlag]bool
}

type RelocationEntry interface {
	VirtualAddress() uintptr
	SymTableIndex() uint16
	RelocationType() RelocationType
}

type SymbolEntry interface {
	Name() string
	Value() []byte
	SectionNo() int16
	StorageClass() StorageClass
	HasAuxEntries() bool
	IsAbsolute() bool
	IsDefined() bool
}

type BinaryFile interface {
	VersionNo() uint16
	Timestamp() uint32
	Flags() map[BinaryFileFlag]bool
	OptionalHeaderSize() uint16
	TargetMagicNumber() BinaryFileTarget
	FileMagicNumber() uint16
	VersionStamp() uint16
	Entrypoint() uintptr
	ExecBase() uintptr
	ExecSize() uint64
	DataBase() uintptr
	DataSize() uint64
	BssSize() uint64
	ImageBase() uintptr
	SectionAlignment() uint32
	FileAlignment() uint32
	SizeOfImage() uint64
	SizeOfHeaders() uint64
	Sections() []BinarySection
	Relocations() []RelocationEntry
	Symbols() []SymbolEntry
}

type BinaryFileBuilder interface {
	WithVersionNo(uint16) *BinaryFileBuilder
	WithTimestamp(uint32) *BinaryFileBuilder
	WithFlags(map[BinaryFileFlag]bool) *BinaryFileBuilder
	WithTargetMagicNumber(uint16) *BinaryFileBuilder
	WithFileMagicNumber(uint16) *BinaryFileBuilder
	WithVersionStamp(uint16) *BinaryFileBuilder
	WithSections([]BinarySection) *BinaryFileBuilder
	WithRelocations([]RelocationEntry) *BinaryFileBuilder
	WithSymbols([]SymbolEntry) *BinaryFileBuilder
}
