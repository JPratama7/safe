package safe

type Valuer interface {
	Unwrap() any
	UnwrapOr(any) any
	Expect(string) any
}
