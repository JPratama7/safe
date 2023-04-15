package safe

import (
	"errors"
	"fmt"
)

type Result[T any] struct {
	err error
	val T
}

func Ok[T any](value T) (res Result[T]) {
	res.val = value
	return
}

func Err[T any](err string) (res Result[T]) {
	res.err = errors.New(err)
	return
}

func AsResult[T any](value T, err error) (res Result[T]) {
	res.val, res.err = value, err
	return
}

func (r *Result[T]) Ok() Option[T] {
	return Some[T](r.val)
}

func (r *Result[T]) Err() Option[error] {
	return Some[error](r.err)
}

func (r *Result[T]) IsOk() (res bool) {
	if r.IsErr() {
		return
	}
	res = NotEmpty(r.val)
	return
}

func (r *Result[T]) IsErr() (res bool) {
	res = r.err != nil
	return
}

func (r *Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(fmt.Errorf("can't unwrap value with err"))
	}
	return r.val
}

func (r Result[T]) Expect(err string) T {
	if r.IsErr() {
		panic(fmt.Errorf(err))
	}
	return r.val
}

func (r Result[T]) UnwrapOr(or T) T {
	if r.IsErr() {
		return or
	}
	return r.val
}
