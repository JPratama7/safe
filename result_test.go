package safe

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestResultOk(t *testing.T) {
	res := result_test_ok()

	assert.Equal(t, res.IsOk(), true)
	assert.Equal(t, res.IsErr(), false)
	assert.NotEmpty(t, res.val)
	assert.NotNil(t, res.val)
}
func TestResultOkSlice(t *testing.T) {
	res := result_test_slices()

	assert.Equal(t, res.IsOk(), true)
	assert.Equal(t, res.IsErr(), false)
	assert.NotEmpty(t, res.val)
	assert.NotNil(t, res.val)
}

func TestResultErr(t *testing.T) {
	res := result_test_none()

	assert.Equal(t, res.IsOk(), false)
	assert.Equal(t, res.IsErr(), true)
	assert.Empty(t, res.val)
}

func result_test_ok() (res Result[TestingWithStruct]) {
	return Ok(TestingWithStruct{
		OuterField:  "croot",
		InnerStruct: InnerStruct{"croot"},
	})
}

func TestMarshalUnmarshalJSONRes(t *testing.T) {
	str := &TestingWithStruct{
		OuterField:  "Hellow World",
		InnerStruct: InnerStruct{"Hellow World World"},
	}
	opt := Option[TestingWithStruct]{Val: str}

	marshal, err := json.Marshal(opt)
	if err != nil {
		t.Fatal(err)
	}

	var opt2 Option[TestingWithStruct]

	err = json.Unmarshal(marshal, &opt2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, opt2.IsSome(), true)
	assert.Equal(t, opt2.IsNone(), false)
	assert.Equal(t, opt2.Unwrap().OuterField, opt.Unwrap().OuterField)
	assert.Equal(t, opt2.Unwrap().InnerStruct.InnerField, opt.Unwrap().InnerStruct.InnerField)
}

func TestMarshalUnmarshalBSONRes(t *testing.T) {
	str := &TestingWithStruct{
		OuterField:  "Hellow World",
		InnerStruct: InnerStruct{"Hellow World World"},
	}
	opt := Option[TestingWithStruct]{Val: str}

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

func result_test_slices() (res Result[[]TestingWithStruct]) {
	return Ok[[]TestingWithStruct]([]TestingWithStruct{{}, {}, {}})
}

func result_test_none() (res Result[int]) {
	return Err[int]("some fancy error message")
}
