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
safe.TestingWithStruct
goos: linux
goarch: amd64
pkg: github.com/JPratama7/safe
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkOkSlicesStruct-8               225667636                5.787 ns/op           0 B/op          0 allocs/op
BenchmarkOkSlicesString-8               181883426                6.864 ns/op           0 B/op          0 allocs/op
BenchmarkOkSlicesInt-8                  194083630                5.786 ns/op           0 B/op          0 allocs/op
BenchmarkOkMapIntString-8               178467765                6.933 ns/op           0 B/op          0 allocs/op
BenchmarkResult_Err-8                   1000000000               0.3706 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8                    1000000000               0.3301 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8                 1000000000               0.3561 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyErr-8             1000000000               0.4281 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyNoErr-8           73274834                14.26 ns/op            0 B/op          0 allocs/op
BenchmarkResult_OkInt-8                 130130389                9.399 ns/op           0 B/op          0 allocs/op
BenchmarkResult_EmptyInt-8              97283092                10.74 ns/op            0 B/op          0 allocs/op
BenchmarkResult_OkString-8              199630009                5.210 ns/op           0 B/op          0 allocs/op
BenchmarkResult_EmptyString-8           234421890                4.800 ns/op           0 B/op          0 allocs/op
BenchmarkOption_Some-8                  1000000000               0.3501 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8                  1000000000               0.3729 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8                362088714                3.142 ns/op           0 B/op          0 allocs/op
BenchmarkOption_IsSome-8                381901064                3.408 ns/op           0 B/op          0 allocs/op
BenchmarkErrorCheck-8                   1000000000               0.3244 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/JPratama7/safe       22.063s

```

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
