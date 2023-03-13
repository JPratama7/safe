package safetypes

import (
	"bytes"
	"github.com/goccy/go-json"
)

type Option[T any] struct {
	Val *T
}

func Some[T any](value T) (o Option[T]) {
	o.Val = &value
	return
}

func None[T any]() (o Option[T]) {
	return
}

func (o *Option[T]) IsSome() (res bool) {
	return o.Val != nil
}

func (o *Option[T]) IsNone() bool {
	return o.Val == nil
}

func (o *Option[T]) Unwrap() T {
	if o.IsNone() {
		return *new(T)
	}
	return *o.Val
}

func (o *Option[T]) UnwrapOr(or T) T {
	if o.IsNone() {
		return or
	}
	return *o.Val
}

func (o *Option[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Val)
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	res := new(T)
	if err := json.Unmarshal(data, res); err != nil {
		if bytes.HasPrefix(data, []byte("{}")) {
			o.Val = nil
			return nil
		}
		return err
	}
	o.Val = res
	return nil
}
