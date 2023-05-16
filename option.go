package safe

import (
	"fmt"
)

type Option[T any] struct {
	val *T
}

func Some[T any](value T) (o Option[T]) {
	o.val = &value
	return
}

func None[T any]() (o Option[T]) {
	return
}

func (o *Option[T]) Some(value T) {
	o.val = &value
}

func (o *Option[T]) None() {
	o.val = new(T)
}

func (o *Option[T]) IsSome() (res bool) {
	res = o.notmissing()
	return
}

func (o *Option[T]) IsNone() (res bool) {
	res = !o.notmissing()
	return
}

func (o Option[T]) Expect(err string) T {
	if o.IsNone() {
		panic(fmt.Errorf(err))
	}
	return *o.val
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("can't unwrap none val")
	}
	return *o.val
}

func (o Option[T]) UnwrapOr(or T) T {
	if o.IsNone() {
		return or
	}
	return *o.val
}

func (o Option[T]) notmissing() (res bool) {
	res = IsNotEmpty(o.val)
	return
}
