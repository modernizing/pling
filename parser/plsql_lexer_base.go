package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type PlSqlLexerBase struct {
	*antlr.BaseLexer
}

func (l *PlSqlLexerBase) IsNewlineAtPos(pos int) bool {
	la := l.GetInputStream().LA(pos)
	return la == -1 || la == '\n'
}
