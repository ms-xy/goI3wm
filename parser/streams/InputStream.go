package streams

import (
	"github.com/ms-xy/logtools"
)

type InputStream interface {
	Next() interface{}
	Peek() interface{}
	Eof() bool
	Croak(message string)
}

func NewInputStream(input <-chan interface{}) InputStream {
	inputStream := &inputStreamImpl{}
	inputStream.init(input)
	return inputStream
}

/**
 * impl
 */

type inputStreamImpl struct {
	input <-chan interface{}
	peek  interface{}
	eof   bool
	pos   int64
}

func (this *inputStreamImpl) init(input <-chan interface{}) {
	this.input = input
	this.eof = false
	this.pos = 0
	this.Next() // read first char into peeked, ignore empty return
}

func (this *inputStreamImpl) Next() interface{} {
	last := this.peek
	if item, ok := <-this.input; ok {
		this.peek = item
		this.pos++
	} else {
		this.eof = true
	}
	return last
}

func (this *inputStreamImpl) Peek() interface{} {
	return this.peek
}

func (this *inputStreamImpl) Eof() bool {
	return this.eof
}

func (this *inputStreamImpl) Croak(message string) {
	logtools.Errorf("%s (%d)", message, this.pos)
}
