package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/inherd/pling/parser"
	"strings"
)

type TreeShapeListener struct {
	*parser.BasePlSqlParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (s *TreeShapeListener) EnterCreate_table(ctx *parser.Create_tableContext) {
	fmt.Println(ctx.GetText())
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {

}

func main() {
	data := `CREATE TABLE employee (
     first_name VARCHAR2(128),
     last_name VARCHAR2(128),
     empID NUMBER,
     salary NUMBER(6) ENCRYPT
);`

	input := antlr.NewInputStream(strings.ToUpper(data))
	lexer := parser.NewPlSqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPlSqlParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Sql_script()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
