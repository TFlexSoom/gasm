package linker

import (
	linker "github.com/tflexsoom/gasm/internal/linker/global"
	"golang.org/x/exp/maps"
)

// https://en.wikipedia.org/wiki/Comparison_of_executable_file_formats
// 1. Support PE
// 2. Support ELF
// 3. Support PE32
// 4. Support Mach-O

type symbolTableImpl struct {
	table        map[linker.ExternRef]uintptr
	tableReverse map[uintptr]linker.ExternRef
	libraries    []linker.Path
	formats      map[string]func([]byte) linker.BinaryFile
}

type SymbolTable interface {
	GetSymbols() []linker.ExternRef
	ReadLibraries([]linker.Path) bool
	IsDefined() bool
	Formats() []string
}

func (symTable symbolTableImpl) addSymbol(ref linker.ExternRef, address uintptr) {
	symTable.table[ref] = address
	symTable.tableReverse[address] = ref
}

func (symTable symbolTableImpl) GetSymbols() []linker.ExternRef {
	return maps.Keys(symTable.table)
}

func (symTable symbolTableImpl) ReadLibraries(paths []linker.Path) bool {
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

func getFormatReaders(formats []string) map[string](func([]byte) linker.BinaryFile) {
	return make(map[string]func([]byte) linker.BinaryFile, 0)
}

func NewTable(formats []string) SymbolTable {
	return symbolTableImpl{
		formats: getFormatReaders(formats),
	}
}
