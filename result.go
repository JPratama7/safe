package safetypes

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
)

type Result[T any] struct {
	err error
	val *T
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		val: &value,
		err: nil,
	}
}

func Err[T any](err string) Result[T] {
	return Result[T]{
		val: nil,
		err: errors.New(err),
	}
}

func AsResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err.Error())
	}
	return Ok(value)
}

func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) Error() error {
	return r.err
}

func (r *Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("can't unwrap err val")
	}
	return *r.val
}

func (r *Result[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.val)
}

func (r *Result[T]) UnmarshalJSON(data []byte) error {
	res := new(T)
	if err := json.Unmarshal(data, res); err != nil {
		if bytes.HasPrefix(data, []byte("{}")) {
			r.val = nil
			return nil
		}
		return err
	}
	r.val = res
	return nil
}

func (r *Result[T]) Value() (driver.Value, error) {
	return fmt.Sprintf("%+v", r.val), nil
}

func (r *Result[T]) Scan(src interface{}) error {
	data, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal val %v %s %s", src, "of type", fmt.Sprintf("%T", src))
	}
	res := new(T)
	if err := json.Unmarshal(data, res); err != nil {
		return err
	}
	r.val = res
	return nil
}
