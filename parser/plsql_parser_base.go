package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type PlSqlParserBase struct {
	*antlr.BaseParser
}

var is_version_12 bool
var is_version_10 bool

func (l *PlSqlParserBase) setVersion12(value bool) {
	is_version_12 = value
}

func (l *PlSqlParserBase) setVersion10(value bool) {
	is_version_10 = value
}

func isVersion12() bool {
	return is_version_12
}

func isVersion10() bool {
	return is_version_10
}
