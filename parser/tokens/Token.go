package tokens

type Token interface {
	Type() string
	Value() interface{}
	String() string
}
