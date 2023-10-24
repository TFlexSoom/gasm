package linker

type ExternRef string
type LibraryName string
type Path string

type BinarySection struct {
	Name    string
	Start   uintptr
	Size    uint64
	Content []byte
	Flags   map[BinarySectionFlag]bool
}

type RelocationEntry struct {
	VirtualAddress uintptr
	SymTableIndex  uint16
	RelocationType RelocationType
}

type SymbolEntry struct {
	Name          string
	Value         []byte
	SectionNo     int16
	StorageClass  StorageClass
	HasAuxEntries bool
	IsAbsolute    bool
	IsDefined     bool
}

type BinaryFileHeader struct {
	Size             uint64
	FileMagicNumber  uint16
	VersionStamp     uint16
	Entrypoint       uintptr
	ImageBase        uintptr
	ImageSize        uint64
	HeaderSize       uint64
	SectionAlignment uint32
	FileAlignment    uint32
}

type BinaryFileSections struct {
	sections    []BinarySection
	execSection *BinarySection
	dataSection *BinarySection
	bssSection  *BinarySection
}

type BinaryFile struct {
	VersionNo         uint16
	Timestamp         uint32
	Flags             map[BinaryFileFlag]bool
	Header            BinaryFileHeader
	TargetMagicNumber BinaryFileTarget
	Sections          BinaryFileSections
	Relocations       []RelocationEntry
	Symbols           []SymbolEntry
}