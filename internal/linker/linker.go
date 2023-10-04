package linker

import "golang.org/x/exp/maps"

// https://en.wikipedia.org/wiki/Comparison_of_executable_file_formats
// 1. Support PE
// 2. Support ELF
// 3. Support PE32
// 4. Support Mach-O

type symbolTableImpl struct {
	table        map[ExternRef]uintptr
	tableReverse map[uintptr]ExternRef
	libraries    map[LibraryName]Path
	formats      map[string]func([]byte) BinaryFile
}

type SymbolTable interface {
	GetSymbols() []ExternRef
	ReadLibraries([]Path) bool
	IsDefined() bool
	Formats() []string
}

func (symTable symbolTableImpl) addSymbol(ref ExternRef, address uintptr) {
	symTable.table[ref] = address
	symTable.tableReverse[address] = ref
}

func (symTable symbolTableImpl) GetSymbols() []ExternRef {
	return maps.Keys(symTable.table)
}

func (symTable symbolTableImpl) ReadLibraries(paths []Path) bool {
	for _, path := range paths {
		// 1. Read File At Path into Respective Executable Struct
		// 2. Read External Symbols Required for the File From Header
		// 3. Read Defined Symbols in the file for byte addresses, Read all into Map
		//  so we don't have to return back to the structure
	}

	return symTable.IsDefined()
}

func (symTable symbolTableImpl) IsDefined() bool {
	return len(symTable.table) == len(symTable.tableReverse)
}

func (symTable symbolTableImpl) Formats() []string {
	return maps.Keys(symTable.formats)
}

func getFormatReaders(formats []string) map[string](func([]byte) BinaryFile) {
	return make(map[string]func([]byte) BinaryFile, 0)
}

func NewTable(formats []string) SymbolTable {
	return symbolTableImpl{
		formats: getFormatReaders(formats),
	}
}
