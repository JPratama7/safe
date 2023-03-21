package safetypes

import (
	"bytes"
	"errors"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson"
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
	if err != nil {
		res.err = err
		return
	}
	res.val = value
	return
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

func (r *Result[T]) MarshalBSON() ([]byte, error) {
	return bson.Marshal(r.val)
}

func (r *Result[T]) UnmarshalJSON(data []byte) error {
	res := new(T)
	if err := json.Unmarshal(data, res); err != nil {
		if bytes.HasPrefix(data, []byte("{}")) {
			r.val = *new(T)
			return nil
		}
		return err
	}
	r.val = *res
	return nil
}

func (r *Result[T]) UnmarshalBSON(data []byte) error {
	res := new(T)
	if err := bson.Unmarshal(data, res); err != nil {
		if bytes.HasPrefix(data, []byte("{}")) {
			r.val = *new(T)
			return nil
		}
		return err
	}
	r.val = *res
	return nil
}
