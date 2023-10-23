package linker

import (
	linker "github.com/tflexsoom/gasm/internal/linker/global"
)

type BinarySectionPE struct {
	name    string
	start   uintptr
	size    uint64
	content []byte
	flags   map[linker.BinarySectionFlag]bool
}

func (binarySection BinarySectionPE) Name() string {
	return binarySection.name
}

func (binarySection BinarySectionPE) Start() uintptr {
	return binarySection.start
}

func (binarySection BinarySectionPE) Size() uint64 {
	return binarySection.size
}

func (binarySection BinarySectionPE) Content() []byte {
	return binarySection.content
}

func (binarySection BinarySectionPE) Flags() map[linker.BinarySectionFlag]bool {
	return binarySection.flags
}

type RelocationEntryPE struct {
	virtualAddress uintptr
	symTableIndex  uint16
	relocationType linker.RelocationType
}

func (relocationEntry RelocationEntryPE) VirtualAddress() uintptr {
	return relocationEntry.virtualAddress
}

func (relocationEntry RelocationEntryPE) SymTableIndex() uint16 {
	return relocationEntry.symTableIndex
}

func (relocationEntry RelocationEntryPE) RelocationType() linker.RelocationType {
	return relocationEntry.relocationType
}

type SymbolEntryPE struct {
	name          string
	value         []byte
	sectionNo     int16
	storageClass  linker.StorageClass
	hasAuxEntries bool
	isAbsolute    bool
	isDefined     bool
}

func (symbolEntry SymbolEntryPE) Name() string {
	return symbolEntry.name
}

func (symbolEntry SymbolEntryPE) Value() []byte {
	return symbolEntry.value
}

func (symbolEntry SymbolEntryPE) SectionNo() int16 {
	return symbolEntry.sectionNo
}

func (symbolEntry SymbolEntryPE) StorageClass() linker.StorageClass {
	return symbolEntry.storageClass
}

func (symbolEntry SymbolEntryPE) HasAuxEntries() bool {
	return symbolEntry.hasAuxEntries
}

func (symbolEntry SymbolEntryPE) IsAbsolute() bool {
	return symbolEntry.isAbsolute
}

func (symbolEntry SymbolEntryPE) IsDefined() bool {
	return symbolEntry.isDefined
}

type OptionalHeaderPE struct {
	Size                  uint16
	MajorOSVersion        uint16
	MinorOSVersion        uint16
	MajorImageVersion     uint16
	MinorImageVersion     uint16
	MajorSubsystemVersion uint16
	MinorSubsystemVersion uint16
	Win32VersionValue     uint64
	Checksum              uint64
	Subsystem             WindowsSubsystem
	DllCharacteristics    []DllCharacteristic
	SizeOfStackReserve    uint64
	SizeOfHeapReserve     uint64
	SizeOfHeapCommit      uint64
	LoaderFlags           uint64
	NumberOfRvaAndSizes   uint64
}

type BinaryFilePE struct {
	versionNo         uint16
	timestamp         uint32
	flags             map[linker.BinaryFileFlag]bool
	optionalHeader    OptionalHeaderPE
	targetMagicNumber linker.BinaryFileTarget

	// 	//-- Optional File Header
	fileMagicNumber uint16
	versionStamp    uint16

	sections    []BinarySectionPE
	execSection *BinarySectionPE
	dataSection *BinarySectionPE
	bssSection  *BinarySectionPE

	imageBase   uintptr
	imageSize   uint64
	headersSize uint64

	sectionAlignment uint32
	fileAlignment    uint32

	entrypoint uintptr

	relocations []RelocationEntryPE
	symbols     []SymbolEntryPE
}

// //-- File Header
func (bfp BinaryFilePE) VersionNo() uint16 {
	return bfp.versionNo
}

func (bfp BinaryFilePE) Timestamp() uint32 {
	return bfp.timestamp
}

func (bfp BinaryFilePE) OptionalHeaderSize() uint16 {
	return bfp.optionalHeader.Size
}

func (bfp BinaryFilePE) Flags() map[linker.BinaryFileFlag]bool {
	return bfp.flags
}

func (bfp BinaryFilePE) TargetMagicNumber() linker.BinaryFileTarget {
	return bfp.targetMagicNumber
}

func (bfp BinaryFilePE) FileMagicNumber() uint16 {
	return bfp.fileMagicNumber
}

func (bfp BinaryFilePE) VersionStamp() uint16 {
	return bfp.versionStamp
}

func (bfp BinaryFilePE) Entrypoint() uintptr {
	return bfp.entrypoint
}

func (bfp BinaryFilePE) ExecBase() uintptr {
	return (*bfp.execSection).Start()
}

func (bfp BinaryFilePE) ExecSize() uint64 {
	return (*bfp.execSection).Size()
}

func (bfp BinaryFilePE) DataBase() uintptr {
	return (*bfp.dataSection).Start()
}

func (bfp BinaryFilePE) DataSize() uint64 {
	return (*bfp.dataSection).Size()
}

func (bfp BinaryFilePE) BssStart() uintptr {
	return (*bfp.bssSection).Start()
}

func (bfp BinaryFilePE) BssSize() uint64 {
	return (*bfp.bssSection).Size()
}

// // -- Windows File Header
func (bfp BinaryFilePE) ImageBase() uintptr {
	return bfp.imageBase
}

func (bfp BinaryFilePE) SectionAlignment() uint32 {
	return bfp.sectionAlignment
}

func (bfp BinaryFilePE) FileAlignment() uint32 {
	return bfp.fileAlignment
}

func (bfp BinaryFilePE) SizeOfImage() uint64 {
	return bfp.imageSize
}

func (bfp BinaryFilePE) SizeOfHeaders() uint64 {
	return bfp.headersSize
}

func (bfp BinaryFilePE) Sections() []linker.BinarySection {
	copy := make([]linker.BinarySection, len(bfp.sections))
	for i, e := range bfp.sections {
		copy[i] = e
	}

	return copy
}

func (bfp BinaryFilePE) Relocations() []linker.RelocationEntry {
	copy := make([]linker.RelocationEntry, len(bfp.relocations))
	for i, e := range bfp.relocations {
		copy[i] = e
	}

	return copy
}

func (bfp BinaryFilePE) Symbols() []linker.SymbolEntry {
	copy := make([]linker.SymbolEntry, len(bfp.symbols))
	for i, e := range bfp.symbols {
		copy[i] = e
	}

	return copy
}

func (bfp BinaryFilePE) WindowsOptionalHeader() OptionalHeaderPE {
	return bfp.optionalHeader
}

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
