// author: Tristan Hilbert
// date: 8/29/2023
// filename: ldatGrammar.go
// desc: Parsing Grammar to Build AST for ldat files
package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// // Lexer
func getAsmLexer() (*lexer.StatefulDefinition, error) {
	return lexer.NewSimple([]lexer.SimpleRule{
		{Name: "COMMENT", Pattern: `;[~\n]*\n`},
		{Name: "SECTION_PREFIX", Pattern: `[.]`},
		{Name: "LABEL_OP", Pattern: `:`},
		{Name: "ARG_SEP", Pattern: `,`},
		{Name: "HEX", Pattern: `0x[\dA-Fa-f]+`},
		{Name: "DECIMAL", Pattern: `-?[\d]+\.[\d]*`},
		{Name: "INTEGER", Pattern: `-?[\d]+`},
		{Name: "QUOTED_VAL", Pattern: `"[~"]*"`},
		{Name: "BOOL_VALUE", Pattern: `true|false`},
		{Name: "IDENTIFIER", Pattern: `[_a-zA-Z][_a-zA-Z\d]*`},
		{Name: "WHITESPACE", Pattern: `[ \t\r\n]+`},
	})
}

func GetAsmParser() (*participle.Parser[Assembly], error) {
	var lexer, err = getAsmLexer()
	if err != nil {
		return nil, err
	}

	return participle.Build[Assembly](
		participle.Lexer(lexer),
		participle.UseLookahead(1),
		participle.Union[Term](
			Comment{},
			Section{},
			Instruction{},
		),
		participle.Union[Argument](
			HexArg{},
			Decimal{},
			Integer{},
			String{},
			Identifier{},
		),
	)
}

//// Grammar

type Assembly struct {
	Pos lexer.Position

	Terms []Term `@@*`
}

type Term interface {
	Term()
	Pos() lexer.Position
}

type Comment struct {
	Position lexer.Position

	Comment string `COMMENT+`
}

func (commentBlock Comment) Term() {}
func (commentBlock Comment) Pos() lexer.Position {
	return commentBlock.Position
}

type Section struct {
	Position lexer.Position

	Name string `section SECTION_PREFIX @IDENTIFIER`
}

func (section Section) Term() {}
func (section Section) Pos() lexer.Position {
	return section.Position
}

type Instruction struct {
	Position lexer.Position

	Label string   `@IDENTIFIER LABEL_OP`
	Name  string   `@IDENTIFIER`
	Args  Argument `(@@ ARG_SEP)* @@?`
}

func (instruction Instruction) Term() {}
func (instruction Instruction) Pos() lexer.Position {
	return instruction.Position
}

type Argument interface {
	Argument()
	Pos() lexer.Position
}

type HexArg struct {
	Position lexer.Position

	Value string `@HEX`
}

func (hexArg HexArg) Argument() {}
func (hexArg HexArg) Pos() lexer.Position {
	return hexArg.Position
}

type Decimal struct {
	Position lexer.Position

	Value string `@DECIMAL`
}

func (decimal Decimal) Argument() {}
func (decimal Decimal) Pos() lexer.Position {
	return decimal.Position
}

type Integer struct {
	Position lexer.Position

	Value string `@INTEGER`
}

func (integer Integer) Argument() {}
func (integer Integer) Pos() lexer.Position {
	return integer.Position
}

type String struct {
	Posotion lexer.Position

	Value string `@QUOTED_VAL`
}

func (str String) Argument() {}
func (str String) Pos() lexer.Position {
	return str.Posotion
}

type Identifier struct {
	Position lexer.Position

	Value string `@IDENTIFIER`
}

func (identifier Identifier) Argument() {}
func (identifier Identifier) Pos() lexer.Position {
	return identifier.Position
}
