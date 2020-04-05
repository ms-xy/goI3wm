package streams

import (
	"github.com/ms-xy/logtools"
)

type CharacterInputStream interface {
	Next() rune
	Peek() rune
	Eof() bool
	Croak(message string)
}

func NewCharacterInputStream(input <-chan rune) CharacterInputStream {
	inputStream := &characterInputStreamImpl{}
	inputStream.init(input)
	return inputStream
}

/**
 * impl
 */

type characterInputStreamImpl struct {
	input <-chan rune
	peek  rune
	eof   bool

	pos  int64
	line int64
	col  int64
}

func (this *characterInputStreamImpl) init(input <-chan rune) {
	this.input = input
	this.eof = false
	this.pos = 0
	this.col = 0
	this.line = 1
	this.Next() // read first char into peeked, ignore empty return
}

func (this *characterInputStreamImpl) Next() rune {
	last := this.peek
	if char, ok := <-this.input; ok {
		this.peek = char
		this.pos++
		if char == '\n' {
			this.line++
			this.col = 0
		} else {
			this.col++
		}
	} else {
		this.eof = true
	}
	return last
}

func (this *characterInputStreamImpl) Peek() rune {
	return this.peek
}

func (this *characterInputStreamImpl) Eof() bool {
	return this.eof
}

func (this *characterInputStreamImpl) Croak(message string) {
	logtools.Errorf("%s (%d:%d)", message, this.line, this.col)
}
