package linker

import (
	"debug/pe"
	"log"
	"os"
)

type ObjectifyOptions struct {
	Files          []string
	OutputLocation string
	Verbose        bool
}

func Objectify(options ObjectifyOptions) error {
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
		peFile, err := pe.Open(file)
		if err != nil {
			return err
		}
		defer peFile.Close()

		// Decide what to do here
		num := 0

		if options.Verbose {
			log.Printf("%d bytes written!", num)
		}
	}

	return nil
}
