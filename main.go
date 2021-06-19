package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/inherd/pling/parser"
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
	input := antlr.NewInputStream("create database hello")
	lexer := parser.NewPlSqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPlSqlParser(stream)
	fmt.Println(p)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//p.BuildParseTrees = true
	//tree := p.Json()
	//antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
