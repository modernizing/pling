#!/usr/bin/env bash

cd parser/g4

antlr -Dlanguage=Go -listener PlSqlLexer.g4 -o ../
antlr -Dlanguage=Go -listener PlSqlParser.g4 -o ../
