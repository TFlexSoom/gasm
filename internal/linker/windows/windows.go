package linker

// type RelocationEntry interface {
// 	VirtualAddress() uintptr
// 	SymTableIndex() uint16
// 	RelocationType() RelocationType
// }

// type SymbolEntry interface {
// 	Name() string
// 	Value() []byte
// 	SectionNo() int16
// 	StorageClass() StorageClass
// 	AuxEntries() bool
// 	IsAbsolute() bool
// 	IsDefined() bool
// }

// .text -> Address of .text section
// .data -> Address of .data section
// etext -> Next Available after .text
// ...

type BinaryFilePE struct {
	versionNo uint16
	timestamp uint32
	flags     map[BinaryFileFlag]bool
	// 	//-- File Header
	// 	VersionNo() uint16
	// 	NumberOfSections() uint16
	// 	Timestamp() uint32
	// 	// symbol start
	// 	NumOfSymbols() uint32
	// 	Flags() map[BinaryFileFlag]bool
	// 	OptionalHeaderSize() uint16
	// 	TargetMagicNumber() BinaryFileTarget

	// 	//-- Optional File Header
	// 	FileMagicNumber() uint16
	// 	VersionStamp() uint16
	// 	ExecBase() uintptr
	// 	ExecSize() uint64
	// 	DataBase() uintptr
	// 	DataSize() uint64
	// 	BssSize() uint64

	// 	// -- Windows File Header
	// 	ImageBase() uintptr
	// 	SectionAlignment() uint32
	// 	FileAlignment() uint32
	// 	SizeOfImage() uint64
	// 	SizeOfHeaders() uint64
	// -- Windows Optional Header
	// MajorOSVersion() uint16
	// MinorOSVersion() uint16
	// MajorImageVersion() uint16
	// MinorImageVersion() uint16
	// MajorSubsystemVersion() uint16
	// MinorSubsystemVersion() uint16
	// Win32VersionValue() uint64
	// Checksum() uint64
	// Subsystem() WindowsSubsystem
	// dllCharacteristics() []DllCharacteristic
	// SizeOfStackReserve() uint64
	// SizeOfHeapReserve() uint64
	// SizeOfHeapCommit() uint64
	// LoaderFlags() uint64
	// NumberOfRvaAndSizes() uint64

	// 	// entryPoint

	// 	// beginning of executable code
	// 	// beginning of initialized data

	// 	// section header structure
	// 	Sections() []BinarySection
	// 	Relocations() []RelocationEntry
	// 	Symbols() []SymbolEntry
}

// type BinaryFile interface {
// 	//-- File Header
// 	VersionNo() uint16
// 	NumberOfSections() uint16
// 	Timestamp() uint32
// 	// symbol start
// 	NumOfSymbols() uint32
// 	Flags() map[BinaryFileFlag]bool
// 	OptionalHeaderSize() uint16
// 	TargetMagicNumber() BinaryFileTarget

// 	//-- Optional File Header
// 	FileMagicNumber() uint16
// 	VersionStamp() uint16
// 	ExecBase() uintptr
// 	ExecSize() uint64
// 	DataBase() uintptr
// 	DataSize() uint64
// 	BssSize() uint64

// 	// -- Windows File Header
// 	ImageBase() uintptr
// 	SectionAlignment() uint32
// 	FileAlignment() uint32
// 	SizeOfImage() uint64
// 	SizeOfHeaders() uint64

// 	// entryPoint

// 	// beginning of executable code
// 	// beginning of initialized data

// 	// section header structure
// 	Sections() []BinarySection
// 	Relocations() []RelocationEntry
// 	Symbols() []SymbolEntry
// }

// type BinaryFileBuilder interface {
// 	WithVersionNo(uint16) *BinaryFileBuilder
// 	WithTimestamp(uint32) *BinaryFileBuilder
// 	WithFlags(map[BinaryFileFlag]bool) *BinaryFileBuilder
// 	WithTargetMagicNumber(uint16) *BinaryFileBuilder
// 	WithFileMagicNumber(uint16) *BinaryFileBuilder
// 	WithVersionStamp(uint16) *BinaryFileBuilder
// 	WithSections([]BinarySection) *BinaryFileBuilder
// 	WithRelocations([]RelocationEntry) *BinaryFileBuilder
// 	WithSymbols([]SymbolEntry) *BinaryFileBuilder
// }
