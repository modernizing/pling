package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"fmt"
	parser "github.com/inherd/pling/parser"
)

type TreeShapeListener struct {
	*parser.BasePlSqlParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewPlSqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewPlSqlParser(stream)
	fmt.Println(p)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//p.BuildParseTrees = true
	//tree := p.Json()
	//antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}