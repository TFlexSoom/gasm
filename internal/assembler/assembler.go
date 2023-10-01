package assembler

type AssemblerOptions struct {
	ProjectLocations []string
	OutputLocation   string
	Language         string
	Verbose          bool
}

func Assemble(options AssemblerOptions) error {
	return nil
}

type ParserOptions struct {
	ProjectLocations []string
	OutputLocation   string
	Verbose          bool
}

func ParseOnly(options ParserOptions) error {
	return nil
}
