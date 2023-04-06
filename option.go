package safe

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/goccy/go-reflect"
	"go.mongodb.org/mongo-driver/bson"
)

type Option[T any] struct {
	val T
}

func Some[T any](value T) (o Option[T]) {
	o.val = value
	return
}

func None[T any]() (o Option[T]) {
	return
}

func (o *Option[T]) Some(value T) {
	o.val = value
}

func (o *Option[T]) None() {
	var val T
	o.val = val
}

func (o *Option[T]) IsSome() (res bool) {
	val := reflect.ValueNoEscapeOf(o.val)
	switch val.Kind() {
	case reflect.Chan, reflect.Slice, reflect.String, reflect.Map, reflect.Array:
		res = val.Len() > 0
	default:
		res = val.IsValid() && !val.IsZero()
	}
	return
}

func (o *Option[T]) IsNone() bool {
	return !o.IsSome()
}

func (o *Option[T]) Unwrap() T {
	if o.IsNone() {
		var val T
		o.val = val
	}
	return o.val
}

func (o *Option[T]) UnwrapOr(or T) T {
	if o.IsNone() {
		return or
	}
	return o.val
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.val)
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	var val T

	if bytes.HasPrefix(data, ByteCheck) {
		o.val = val
		return nil
	}

	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	o.val = val
	return nil
}

func (o Option[T]) MarshalBSON() ([]byte, error) {
	return bson.Marshal(o.val)
}

func (o *Option[T]) UnmarshalBSON(data []byte) error {
	var val T

	if bytes.Equal(data, []byte{}) {
		o.val = val
		return nil
	}

	if err := bson.Unmarshal(data, &val); err != nil {
		return err
	}
	o.val = val
	return nil
}
