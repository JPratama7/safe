package safe

import (
	"errors"
	"testing"
)

type InnerStruct struct {
	InnerField string `json:"innerField" bson:"innerField"`
}

type TestingWithStruct struct {
	OuterField  string `bson:"outerField" json:"outerField"`
	InnerStruct `bson:"innerStruct" json:"innerStruct"`
}

func BenchmarkOkSlicesStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := Ok([]TestingWithStruct{
			{
				OuterField:  "croot",
				InnerStruct: InnerStruct{"croot"},
			},
			{
				OuterField:  "croot",
				InnerStruct: InnerStruct{"croot"},
			},
		})
		val.IsOk()
	}
	b.ReportAllocs()
}
func BenchmarkOkSlicesString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := Ok([]string{"", "", ""})
		val.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkOkSlicesInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := Ok([]int{0, 0, 0})
		val.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_Err(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Err[int]("some fancy error message")
	}
	b.ReportAllocs()
}

func BenchmarkResult_Ok(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
	}
	b.ReportAllocs()
}

func BenchmarkResultTestOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result_test_ok()
	}
	b.ReportAllocs()
}

func BenchmarkAsResultEmptyErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := AsResult[TestingWithStruct](emptyStructErr())
		val.IsErr()
	}
	b.ReportAllocs()
}
func BenchmarkAsResultEmptyNoErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := AsResult[TestingWithStruct](emptyStruct())
		val.IsOk()
		val.Unwrap()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok(23)
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkIntZeroVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok(23)
		res.IsOkZeroVal()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkIntOTF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Result[int]{val: 23}
		res.IsOkOTFReflect()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok(0)
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyIntZeroVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok(0)
		res.IsOkZeroVal()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyIntOTF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Result[int]{}
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok("hello world")
		res.IsOkOTFReflect()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkStringZeroVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok("hello world")
		res.IsOkZeroVal()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkStringOTF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Result[string]{
			val: "hello world",
		}
		res.IsOkOTFReflect()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok("")
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyStringZeroVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok("")
		res.IsOkZeroVal()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyStringOTF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Result[string]{}
		res.IsOkOTFReflect()
	}
	b.ReportAllocs()
}

func BenchmarkOption_Some(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Some[TestingWithStruct](TestingWithStruct{})
		Some[TestingWithStruct](TestingWithStruct{})
		Some[TestingWithStruct](TestingWithStruct{})
	}
	b.ReportAllocs()
}

func BenchmarkOption_None(b *testing.B) {
	for i := 0; i < b.N; i++ {
		None[int]()
	}
	b.ReportAllocs()
}

func BenchmarkOption_IsNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := None[TestingWithStruct]()
		val.IsNone()
	}
	b.ReportAllocs()
}

func BenchmarkOption_IsSome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := Some[TestingWithStruct](TestingWithStruct{})
		val.IsSome()
	}
	b.ReportAllocs()
}

func BenchmarkErrorCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := AsResult(TestingWithStruct{}, errors.New("some fancy error message"))
		err.IsErr()

	}
	b.ReportAllocs()
}

func emptyStructErr() (data TestingWithStruct, err error) {
	err = errors.New("some fancy error message")
	return
}

func emptyStruct() (data TestingWithStruct, err error) {
	data = TestingWithStruct{
		OuterField:  "testing",
		InnerStruct: InnerStruct{InnerField: "testing2"},
	}
	return
}
