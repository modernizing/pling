# PL/SQLing

> build domain model from PL/SQL file.

## TODO

 - [x] import parser & grammar
 - [ ] analyse
    - [ ] create table
    - [ ] create index

## Dev

1. compile antlr

```
./script/compile-antlr.sh
```

2. fix generate code typo

```golang
type PlSqlLexer struct {
	*PlSqlLexerBase  // change to PlSqlLexerBase
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}
```

License
---

This code is distributed under the MPL license. See `LICENSE` in this directory.

