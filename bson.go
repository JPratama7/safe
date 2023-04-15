package safe

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson"
)

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
