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
go test --bench=.
goos: linux
goarch: amd64
pkg: github.com/JPratama7/safe
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkResult_Err-8           1000000000               0.3575 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8            1000000000               0.3289 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8         1000000000               0.2877 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8          1000000000               0.2966 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8          1000000000               0.3057 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8        1000000000               0.3048 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsSome-8        1000000000               0.2941 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/JPratama7/safe       2.452s
```

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
