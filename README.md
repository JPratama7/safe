# safetypes
Rust like result and option implementation for golang
## Examples

```go
import "github.com/JPratama7/safe"
```

### Option
```go
import "github.com/JPratama7/safe"

func checkUnwrap(opt safe.Option[int]) {
    if opt.IsSome() {
        println(opt.Unwrap())
    } else {
        panic("poor option :(")
    }
}
```
```go
import "github.com/JPratama7/safe"

func checkUnwrapOr(opt safe.Option[int]) {
    println(opt.UnwrapOr(10))
}
```
```go
import "github.com/JPratama7/safe"

func retrunOption(some bool) (opt safe.Option[int]) {
    if some {
        return opt.Some(7)
    }
    return opt.None()
}
```
```go
import "github.com/JPratama7/safe"

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
import "github.com/JPratama7/safe"

func retrunResult(some bool) (res safe.Result[int]) {
    if some {
        return res.Ok(7)
    }
    return res.Err("some fancy error msg")
}
```

### Benchmark
```bash
go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/JPratama7/safe
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkOkSlices-8                      7604680               153.7 ns/op            88 B/op          2 allocs/op
BenchmarkResult_Err-8                   1000000000               0.4449 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8                    1000000000               0.4239 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8                 1000000000               0.4808 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyErr-8             1000000000               0.4364 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyNoErr-8           10902249               103.0 ns/op            32 B/op          1 allocs/op
BenchmarkOption_Some-8                  1000000000               0.4301 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8                  1000000000               0.4348 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8                1000000000               0.4260 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsSome-8                1000000000               0.4255 ns/op          0 B/op          0 allocs/op
BenchmarkErrorCheck-8                   1000000000               0.4394 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/JPratama7/safe       6.976s
```

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
