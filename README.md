# safetypes
Rust like result and option implementation for golang

## Examples

```go
import safe "github.com/JPratama7/safe"
```

### Option
```go
func checkUnwrap(opt safe.Option[int]) {
    if opt.IsSome() {
        println(opt.Unwrap())
    } else {
        panic("poor option :(")
    }
}
```
```go
func checkUnwrapOr(opt safe.Option[int]) {
    println(opt.UnwrapOr(10))
}
```
```go
func retrunOption(some bool) (opt safe.Option[int]) {
    if some {
        return opt.Some(7)
    }
    return opt.None()
}
```
```go
type Test struct {
    Field safe.Option[int]
}

func jsonMarshal(t Test) {
    res := safe.AsResult(json.Marshal(s))
    if res.IsOk() {
        // if some: "Test{Field: 7}"
        // if none: "Test{Field: {}}"
        println(res.Unwrap())
    } else {
        panic(res.Error())
    }
}
```

### Result
```go
func checkUnwrap(res safe.Result[int]) {
    if res.IsOk() {
        println(res.Unwrap())
    } else {
        panic(res.Error())
    }
}
```
```go
func retrunResult(some bool) (res safe.Result[int]) {
    if some {
        return res.Ok(7)
    }
    return res.Err("some fancy error msg")
}
```

### Benchmark
```
go test -bench=. -count 5 -run=.
goos: linux
goarch: amd64
pkg: github.com/JPratama7/safetypes
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkResult_Err-8           1000000000               0.2972 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8           1000000000               0.3000 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8           1000000000               0.3105 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8           1000000000               0.2820 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8           1000000000               0.3020 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8            1000000000               0.3029 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8            1000000000               0.3049 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8            1000000000               0.2903 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8            1000000000               0.3135 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8            1000000000               0.3055 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8         1000000000               0.3090 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8         1000000000               0.3035 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8         1000000000               0.3025 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8         1000000000               0.3206 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8         1000000000               0.3259 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8          1000000000               0.3341 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8          1000000000               0.2934 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8          1000000000               0.3058 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8          1000000000               0.3061 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8          1000000000               0.2949 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8          1000000000               0.3153 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8          1000000000               0.3149 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8          1000000000               0.3164 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8          1000000000               0.3085 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8          1000000000               0.3450 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8        1000000000               0.3386 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8        1000000000               0.3206 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8        1000000000               0.3958 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8        1000000000               0.5219 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8        1000000000               0.4973 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsSome-8        1000000000               0.3281 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsSome-8        1000000000               0.3509 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsSome-8        1000000000               0.3150 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsSome-8        1000000000               0.3455 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsSome-8        1000000000               0.3798 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/JPratama7/safetypes  13.001s

```

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
