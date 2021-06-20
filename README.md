# PL/SQLing

> build domain model from PL/SQL file.

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

@ 2021 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MPL license. See `LICENSE` in this directory.

