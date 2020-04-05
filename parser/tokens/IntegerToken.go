package tokens

import (
	"fmt"
)

type IntegerToken struct {
	Token
	value int
}

func NewIntegerToken(value int) *IntegerToken {
	return &IntegerToken{
		value: value,
	}
}

func (this *IntegerToken) Type() string {
	return "int"
}

func (this *IntegerToken) Value() interface{} {
	return this.value
}

func (this *IntegerToken) String() string {
	return fmt.Sprintf("<IntegerToken value='%d'>", this.value)
}
