package safe

import (
	"bytes"
	"errors"
	"github.com/goccy/go-json"
	goref "github.com/goccy/go-reflect"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type Result[T any] struct {
	err    error
	val    T
	refVal reflect.Value
}

func Ok[T any](value T) (res Result[T]) {
	res.val = value
	res.refVal = reflect.ValueOf(value)
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
	switch r.refVal.Kind() {
	case reflect.Chan, reflect.Slice, reflect.Map, reflect.Array:
		res = r.refVal.Len() > 0
		break
	default:
		res = r.refVal.IsValid() && !r.refVal.IsZero()
		break
	}
	return
}

func (r *Result[T]) IsOkOTFReflect() (res bool) {
	val := reflect.ValueOf(r.val)
	if r.IsErr() {
		return
	}

	switch val.Kind() {
	case reflect.Chan, reflect.Slice, reflect.Map:
		res = !val.IsNil()
		return
	case reflect.Array, reflect.Struct:
		res = !val.Equal(reflect.Zero(val.Type()))
	default:
		res = val.IsValid() && !val.IsZero()
		return
	}
	return
}

func (r *Result[T]) IsOkGoReflect() (res bool) {
	if r.IsErr() {
		return
	}
	res = Checker(goref.ValueNoEscapeOf(r.val))
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

func (r *Result[T]) Expect(err error) T {
	if r.IsErr() {
		panic(err)
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
	r.refVal = reflect.ValueOf(res)
	return nil
}

func (r Result[T]) MarshalBSON() ([]byte, error) {
	return bson.Marshal(r.val)
}

func (r *Result[T]) UnmarshalBSON(data []byte) error {
	res := new(T)

	if bytes.Equal(data, []byte{}) {
		r.val = *res
		return nil
	}

	if err := bson.Unmarshal(data, res); err != nil {
		return err
	}
	r.val = *res
	r.refVal = reflect.ValueOf(res)
	return nil
}
