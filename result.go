package safe

import (
	"bytes"
	"errors"
	"github.com/goccy/go-json"
	"github.com/goccy/go-reflect"
)

type Result[T any] struct {
	err error
	val T `bson:"val"`
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

func (r *Result[T]) Ok(value T) {
	r.val = value
}

func (r *Result[T]) Err(err string) {
	r.err = errors.New(err)
}

func (r *Result[T]) IsOk() (res bool) {
	val := reflect.ValueNoEscapeOf(r.val)
	if r.IsErr() {
		return
	}

	switch val.Kind() {
	case reflect.Chan, reflect.Slice, reflect.String, reflect.Map, reflect.Array:
		res = val.Len() > 0
	default:
		res = val.IsValid() && !val.IsZero()
	}
	return
}

func (r *Result[T]) IsErr() (res bool) {
	res = r.err != nil
	return
}

func (r *Result[T]) Error() error {
	return r.err
}

func (r *Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("can't unwrap err val")
	}
	return r.val
}

func (r *Result[T]) UnwrapOr(or T) T {
	if r.IsOk() {
		return or
	}
	return r.val
}

func (r Result[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.val)
}

func (r *Result[T]) UnmarshalJSON(data []byte) error {
	res := new(T)

	if bytes.HasPrefix(data, ByteCheck) {
		r.val = *new(T)
		return nil
	}

	if err := json.Unmarshal(data, res); err != nil {

		return err
	}

	r.val = *res
	return nil
}
