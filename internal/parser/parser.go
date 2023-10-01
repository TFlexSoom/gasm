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
		{Name: "WHITESPACE", Pattern: `[ \t]+`},
		{Name: "EOL", Pattern: `\r?\n`},
		{Name: "COMMENT_OP", Pattern: `;`},
		{Name: "SECTION_PREFIX", Pattern: `.`},
		{Name: "LABEL_OP", Pattern: `:`},
		{Name: "ARG_SEP", Pattern: `,`},
		{Name: "HEX", Pattern: `0x[\d]+`},
		{Name: "DECIMAL", Pattern: `[\d]+\.[\d]*`},
		{Name: "INTEGER", Pattern: `[\d]+`},
		{Name: "QUOTED_VAL", Pattern: `"[~"]*"`},
		{Name: "BOOL_VALUE", Pattern: `true|false`},
		{Name: "IDENTIFIER", Pattern: `[a-zA-Z][a-zA-Z\d]*`},
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
	)
}

//// Grammar

type Assembly struct {
	Pos lexer.Position
}
