package main

import (
	"github.com/ms-xy/goI3wm/parser"
	"github.com/ms-xy/logtools"
	"os"
)

func main() {
	input := make(chan string)
	go func() {
		for _, arg := range os.Args[1:] {
			input <- arg
		}
		close(input)
	}()
	lexer := parser.NewStringLexer(input)
	for !lexer.Eof() {
		token := lexer.Next()
		logtools.Info(token.String())
	}
}
