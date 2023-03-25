package safe

import (
	"errors"
	"testing"
)

type InnerStruct struct {
	InnerField string
}

type TestingWithStruct struct {
	OuterField string
	InnerStruct
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

func BenchmarkOkSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := Ok([]TestingWithStruct{
			TestingWithStruct{
				OuterField:  "croot",
				InnerStruct: InnerStruct{"croot"},
			},
			TestingWithStruct{
				OuterField:  "croot",
				InnerStruct: InnerStruct{"croot"},
			},
		})
		val.IsOk()
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
