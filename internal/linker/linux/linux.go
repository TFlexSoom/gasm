package linker

type OptionalHeaderELF struct {
	ident     string
	typeElf   E_TYPE
	machine   E_MACHINE
	version   E_VERSION
	entry     uintptr
	phoff     uint64
	shoff     uint64
	flags     uint32
	ehsize    uint16
	phentsize uint16
	phnum     uint16
	shentsize uint16
	shnum     uint16
	shstrndx  uint16
}
