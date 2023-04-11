package safe

import (
	"encoding/json"
	"github.com/goccy/go-reflect"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestSome(t *testing.T) {
	res := option_test_some()

	assert.Equal(t, res.IsSome(), true)
	assert.Equal(t, res.IsNone(), false)
	assert.NotEmpty(t, res.val)
}

func TestSomeStruct(t *testing.T) {
	res := Some(TestingWithStruct{OuterField: "Test", InnerStruct: InnerStruct{InnerField: "Test"}})

	assert.Equal(t, res.IsSome(), true)
	assert.Equal(t, res.IsNone(), false)
	assert.NotEmpty(t, res.val)
}

func TestNone(t *testing.T) {
	res := option_test_none()

	assert.Equal(t, res.IsSome(), false)
	assert.Equal(t, res.IsNone(), true)
	assert.Empty(t, res.val)
}

func TestNoneStruct(t *testing.T) {
	res := Some(TestingWithStruct{})
	assert.Equal(t, res.IsSome(), false)
	assert.Equal(t, res.IsNone(), true)
	assert.Empty(t, res.val)
}

func TestMarshalUnmarshalJSONOpt(t *testing.T) {
	str := &TestingWithStruct{
		OuterField:  "Hellow World",
		InnerStruct: InnerStruct{"Hellow World World"},
	}
	opt := Some(str)

	marshal, err := json.Marshal(opt)
	if err != nil {
		t.Fatal(err)
	}

	result := new(TestingWithStruct)
	opt2 := Some(result)

	err = json.Unmarshal(marshal, &opt2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, opt2.IsSome(), true)
	assert.Equal(t, opt2.IsNone(), false)
	assert.Equal(t, reflect.DeepEqual(opt2.Unwrap(), opt.Unwrap()), true)
	assert.Equal(t, opt2.Unwrap().OuterField, opt.Unwrap().OuterField)
	assert.Equal(t, opt2.Unwrap().InnerStruct.InnerField, opt.Unwrap().InnerStruct.InnerField)
}

func TestMarshalUnmarshalBSONOpt(t *testing.T) {
	str := &TestingWithStruct{
		OuterField:  "Hellow World",
		InnerStruct: InnerStruct{"Hellow World World"},
	}
	opt := Some(str)

	marshal, err := bson.Marshal(opt)
	if err != nil {
		t.Fatal(err)
	}

	var opt2 Option[TestingWithStruct]

	err = bson.Unmarshal(marshal, &opt2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, opt2.IsSome(), true)
	assert.Equal(t, opt2.IsNone(), false)
	assert.Equal(t, opt2.Unwrap().OuterField, opt.Unwrap().OuterField)
	assert.Equal(t, opt2.Unwrap().InnerStruct.InnerField, opt.Unwrap().InnerStruct.InnerField)
}

func option_test_some() (opt Option[int]) {
	return Some[int](7)
}

func option_test_none() (opt Option[int]) {
	return None[int]()
}
