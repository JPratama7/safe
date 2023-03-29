package safe

import (
	"bytes"
	"github.com/goccy/go-json"
)

type Option[T any] struct {
	Val *T `bson:"val"`
}

func Some[T any](value T) (o Option[T]) {
	o.Val = &value
	return
}

func None[T any]() (o Option[T]) {
	return
}

func (o *Option[T]) Some(value T) {
	o.Val = &value
}

func (o *Option[T]) None() {
	o.Val = nil
}

func (o *Option[T]) IsSome() (res bool) {
	return o.Val != nil
}

func (o *Option[T]) IsNone() bool {
	return !o.IsSome()
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

func (o Option[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Val)
}

//func (o *Option[T]) MarshalBSON() ([]byte, error) {
//	return bson.Marshal(o.Val)
//}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	res := new(T)

	if bytes.HasPrefix(data, ByteCheck) {
		o.Val = res
		return nil
	}

	if err := json.Unmarshal(data, res); err != nil {
		return err
	}
	o.Val = res
	return nil
}

//func (o *Option[T]) UnmarshalBSON(data []byte) error {
//	res := new(T)
//
//	fmt.Printf("data: %s\n", string(data))
//
//	if bytes.HasPrefix(data, ByteCheck) {
//		o.Val = new(T)
//		return nil
//	}
//
//	if err := bson.Unmarshal(data, res); err != nil {
//		return err
//	}
//	o.Val = res
//	return nil
//}
