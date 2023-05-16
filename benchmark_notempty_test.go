package safe

import "testing"

func BenchmarkNotEmptyNotEmptyInt(b *testing.B) {
	data := 10
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}

func BenchmarkNotEmptyEmptyInt(b *testing.B) {
	data := 0
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}

func BenchmarkNotEmptyEmptyString(b *testing.B) {
	data := ""
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}

func BenchmarkNotEmptyNotEmptyString(b *testing.B) {
	data := "testing"
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}

func BenchmarkNotEmptyNotEmptyStructValue(b *testing.B) {
	data := TestingWithStruct{OuterField: "testing",
		InnerStruct: InnerStruct{InnerField: "testing2"}}
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}

func BenchmarkNotEmptyNotEmptyStructPtr(b *testing.B) {
	data := &TestingWithStruct{OuterField: "testing",
		InnerStruct: InnerStruct{InnerField: "testing2"}}
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}

func BenchmarkNotEmptyEmptyStructPtr(b *testing.B) {
	data := &TestingWithStruct{}
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}

func BenchmarkNotEmptyEmptyStructValue(b *testing.B) {
	data := TestingWithStruct{}
	for i := 0; i < b.N; i++ {
		IsNotEmpty(data)
	}
	b.ReportAllocs()
}
