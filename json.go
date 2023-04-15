package safe

import (
	"bytes"
	"github.com/goccy/go-json"
)

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
