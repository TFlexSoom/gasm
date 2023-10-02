// author: Tristan Hilbert
// date: 10/02/2023
// filename: parser.go
// desc: Parsing Grammar to Build AST for assembly files of all languages
//
//	Currently supporting:
//	- x86
package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// // Lexer
func getAsmLexer() (*lexer.StatefulDefinition, error) {
	return lexer.New(lexer.Rules{
		"AlwaysPre": {
			{Name: "COMMENT", Pattern: `;[^\n]*\n`, Action: nil},
		},
		"AlwaysPost": {
			{Name: "HEX", Pattern: `0x[\dA-Fa-f]+`, Action: nil},
			{Name: "DECIMAL", Pattern: `-?[\d]+\.[\d]*`, Action: nil},
			{Name: "INTEGER", Pattern: `-?[\d]+`, Action: nil},
			{Name: "SNGL_QUOTED_VAL", Pattern: `'[^']*'`, Action: nil},
			{Name: "DBL_QUOTED_VAL", Pattern: `"[^"]*"`, Action: nil},
			{Name: "BOOL_VALUE", Pattern: `true|false`, Action: nil},
			{Name: "IDENTIFIER", Pattern: `[_a-zA-Z][_@a-zA-Z\d]*`, Action: nil},
			{Name: "WHITESPACE", Pattern: `[ \t\r\n]+`, Action: nil},
		},
		"StateStarts": {
			{Name: "SUB_ADDR_BEGIN", Pattern: `\[`, Action: lexer.Push("Address")},
			{Name: "SUB_EXPR_BEGIN", Pattern: `\(`, Action: lexer.Push("Expression")},
		},
		"Root": {
			lexer.Include("AlwaysPre"),
			{Name: "SECTION_PREFIX", Pattern: `\.`, Action: nil},
			{Name: "LABEL_OP", Pattern: `:`, Action: nil},
			{Name: "ARG_SEP", Pattern: `,`, Action: nil},
			lexer.Include("StateStarts"),
			lexer.Include("AlwaysPost"),
		},
		"Constexpr": {
			{Name: "OPERATOR", Pattern: `[+\-*/]`, Action: nil},
		},
		"Address": {
			lexer.Include("AlwaysPre"),
			{Name: "SUB_ADDR_END", Pattern: `\]`, Action: lexer.Pop()},
			lexer.Include("StateStarts"),
			lexer.Include("Constexpr"),
			lexer.Include("AlwaysPost"),
		},
		"Expression": {
			lexer.Include("AlwaysPre"),
			{Name: "SUB_EXPR_END", Pattern: `\)`, Action: lexer.Pop()},
			lexer.Include("StateStarts"),
			lexer.Include("Constexpr"),
			lexer.Include("AlwaysPost"),
		},
	})
}

func GetAsmParser() (*participle.Parser[Assembly], error) {
	var lexer, err = getAsmLexer()
	if err != nil {
		return nil, err
	}

	return participle.Build[Assembly](
		participle.Lexer(lexer),
		participle.UseLookahead(2),
		participle.Elide("WHITESPACE", "COMMENT"),
		participle.Union[Term](
			Section{},
			Instruction{},
			Label{},
		),
		participle.Union[Argument](
			HexArg{},
			Decimal{},
			Integer{},
			String{},
			Identifier{},
			SubAddress{},
			SubExpression{},
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

type Expression interface {
	Expression()
	Pos() lexer.Position
}

type Section struct {
	Position lexer.Position

	Name string `"section" SECTION_PREFIX @IDENTIFIER (?! LABEL_OP)`
}

func (section Section) Term() {}
func (section Section) Pos() lexer.Position {
	return section.Position
}

type Instruction struct {
	Position lexer.Position

	Prefix *Label   `@@? (?= IDENTIFIER)`
	Name   string   `@IDENTIFIER (?! LABEL_OP)`
	Args   Argument `@@? (ARG_SEP @@)*`
}

func (instruction Instruction) Term() {}
func (instruction Instruction) Pos() lexer.Position {
	return instruction.Position
}

type Label struct {
	Position lexer.Position

	Name string `@IDENTIFIER LABEL_OP`
}

func (label Label) Term() {}
func (label Label) Pos() lexer.Position {
	return label.Position
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

	Value string `(@SNGL_QUOTED_VAL | @DBL_QUOTED_VAL)`
}

func (str String) Argument() {}
func (str String) Pos() lexer.Position {
	return str.Posotion
}

type Identifier struct {
	Position lexer.Position

	Value string `@IDENTIFIER (?! LABEL_OP)`
}

func (identifier Identifier) Argument() {}
func (identifier Identifier) Pos() lexer.Position {
	return identifier.Position
}

type SubExpression struct {
	Position lexer.Position

	Sub        Argument   `SUB_EXPR_BEGIN @@`
	Operations []Operator `@@* SUB_EXPR_END`
}

func (subExpression SubExpression) Argument() {}
func (subExpression SubExpression) Pos() lexer.Position {
	return subExpression.Position
}

type SubAddress struct {
	Position lexer.Position

	Sub        Argument   `SUB_ADDR_BEGIN @@`
	Operations []Operator `@@* SUB_ADDR_END`
}

func (subAddress SubAddress) Argument() {}
func (subAddress SubAddress) Pos() lexer.Position {
	return subAddress.Position
}

type Operator struct {
	Position lexer.Position

	Symbol  string   `@OPERATOR`
	Operand Argument `@@`
}
