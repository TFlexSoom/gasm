package assembler

import (
	"debug/pe"
	"log"
	"os"

	"github.com/alecthomas/repr"
	"github.com/tflexsoom/gasm/internal/parser"
)

type AssemblerOptions struct {
	Files          []string
	Includes       []string
	OutputLocation string
	Language       string
	Verbose        bool
}

func Assemble(options AssemblerOptions) error {
	asmParser, err := parser.GetAsmParser()
	if err != nil {
		return err
	}

	for _, file := range options.Files {
		reader, err := os.Open(file)
		if err != nil {
			return err
		}
		defer reader.Close()

		ast, err := asmParser.Parse(file, reader)
		if err != nil {
			return err
		}

	}

	for _, file := range options.Files {
		peFile, err := pe.Open(file)
		if err != nil {
			return err
		}
		defer peFile.Close()

	}

	outFile, err := os.OpenFile(
		options.OutputLocation,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return nil
}

type ParserOptions struct {
	Files          []string
	OutputLocation string
	Verbose        bool
}

func ParseOnly(options ParserOptions) error {
	asmParser, err := parser.GetAsmParser()
	if err != nil {
		return err
	}

	outFile, err := os.OpenFile(
		options.OutputLocation,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for _, file := range options.Files {
		reader, err := os.Open(file)
		if err != nil {
			return err
		}
		defer reader.Close()

		ast, err := asmParser.Parse(file, reader)
		if err != nil {
			return err
		}

		num, err := outFile.WriteString(repr.String(ast))
		if err != nil {
			return err
		}

		if options.Verbose {
			log.Printf("%d bytes written!", num)
		}
	}

	return nil
}
