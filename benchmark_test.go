package safetypes

import "testing"

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
	}
	b.ReportAllocs()
}

func BenchmarkResultTestOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result_test_ok()
	}
	b.ReportAllocs()
}

func BenchmarkOption_Some(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Some[int](7)
		Some[int](8)
		Some[int](9)
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
