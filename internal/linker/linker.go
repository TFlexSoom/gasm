package linker

import (
	"log"
	"os"

	"github.com/alecthomas/repr"
	global "github.com/tflexsoom/gasm/internal/linker/global"
	windows "github.com/tflexsoom/gasm/internal/linker/windows"
	"golang.org/x/exp/maps"
)

// https://en.wikipedia.org/wiki/Comparison_of_executable_file_formats
// 1. Support PE
// 2. Support ELF
// 3. Support PE32
// 4. Support Mach-O

type symbolTableImpl struct {
	table        map[global.ExternRef]uintptr
	tableReverse map[uintptr]global.ExternRef
	// libraries    []global.Path
	formats map[string]global.BinaryFileReader
}

func (symTable symbolTableImpl) AddSymbol(ref global.ExternRef, address uintptr) {
	symTable.table[ref] = address
	symTable.tableReverse[address] = ref
}

func (symTable symbolTableImpl) GetSymbols() []global.ExternRef {
	return maps.Keys(symTable.table)
}

func (symTable symbolTableImpl) ReadLibraries(paths []global.Path) bool {
	// for _, path := range paths {
	// 1. Read File At Path into Respective Executable Struct
	// 2. Read External Symbols Required for the File From Header
	// 3. Read Defined Symbols in the file for byte addresses, Read all into Map
	//  so we don't have to return back to the structure
	// }

	return symTable.IsDefined()
}

func (symTable symbolTableImpl) IsDefined() bool {
	return len(symTable.table) == len(symTable.tableReverse)
}

func (symTable symbolTableImpl) Formats() []string {
	return maps.Keys(symTable.formats)
}

func getFormatReaders(formats []string) map[string]global.BinaryFileReader {
	return map[string]global.BinaryFileReader{
		"COFF": global.GetCOFFReader(),
		"PE":   windows.GetPEReader(),
		// "ELF": linux.GetELFReader(),
	}
}

func NewTable(formats []string) symbolTableImpl {
	return symbolTableImpl{
		formats: getFormatReaders(formats),
	}
}

type ObjectifyOptions struct {
	Files          []string
	OutputLocation string
	Verbose        bool
}

func Objectify(options ObjectifyOptions) error {
	peReader := windows.GetPEReader()

	outFile, err := os.OpenFile(
		options.OutputLocation,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}

	for _, file := range options.Files {
		reader, err := os.Open(file)
		if err != nil {
			return err
		}

		result, err := peReader(reader)
		if err != nil {
			return err
		}

		num, err := outFile.WriteString(repr.String(result))
		if err != nil {
			return err
		}

		if options.Verbose {
			log.Printf("%d bytes written!", num)
		}
	}

	return nil
}
