package parser

import (
	"github.com/ms-xy/goI3wm/parser/streams"
	"github.com/ms-xy/goI3wm/parser/tokens"
	"regexp"
	"strconv"
)

type Lexer interface {
	Next() tokens.Token
	Peek() tokens.Token
	Eof() bool
	Croak(message string)
}

func NewStringLexer(input <-chan string) Lexer {
	lexer := &stringLexer{}
	lexer.init(input)
	return lexer
}

/**
 * StringLexer
 */

type stringLexer struct {
	input streams.InputStream
	peek  tokens.Token
	eof   bool
}

var (
	keywords = []string{
		"move", "to",
	}
	re_number *regexp.Regexp
)

func init() {
	re_number = regexp.MustCompile("^\\d+$")
}

func (this *stringLexer) init(input <-chan string) {
	translatedInput := make(chan interface{})
	go func() {
		for str := range input {
			translatedInput <- str
		}
		close(translatedInput)
	}()
	this.input = streams.NewInputStream(translatedInput)
	this.eof = false
	this.Next()
}

func (this *stringLexer) Next() tokens.Token {
	current := this.peek
	this.peek = this.read_next()
	return current
}

func (this *stringLexer) Peek() tokens.Token {
	return this.peek
}

func (this *stringLexer) Eof() bool {
	return this.eof
}

func (this *stringLexer) Croak(message string) {
	this.input.Croak(message)
}

func (this *stringLexer) read_next() tokens.Token {
	if !this.input.Eof() {
		str, _ := (this.input.Next()).(string)
		if re_number.MatchString(str) {
			return this.read_integer(str)
		}
	}
	this.eof = true
	return nil
}

func (this *stringLexer) read_integer(str string) *tokens.IntegerToken {
	if number, err := strconv.Atoi(str); err != nil {
		panic(err)
	} else {
		return tokens.NewIntegerToken(number)
	}
}
