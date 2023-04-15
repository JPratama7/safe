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
go test -bench=. -count=5
goos: linux
goarch: amd64
pkg: github.com/JPratama7/safe
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkOkSlicesStruct-8               14212250                90.36 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesStruct-8               14285074                86.78 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesStruct-8               13850388                87.35 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesStruct-8               14020578                88.04 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesStruct-8               12869972                89.48 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesString-8               14812968                79.81 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesString-8               13490724                81.34 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesString-8               13533954                81.55 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesString-8               13508130                81.08 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesString-8               15963933                84.87 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesInt-8                  15223528                81.45 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesInt-8                  15077181                82.55 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesInt-8                  14913943                80.27 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesInt-8                  13870224                79.43 ns/op           24 B/op          1 allocs/op
BenchmarkOkSlicesInt-8                  15237476                81.07 ns/op           24 B/op          1 allocs/op
BenchmarkResult_Err-8                   1000000000               0.3375 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8                   1000000000               0.3291 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8                   1000000000               0.3357 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8                   1000000000               0.3278 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Err-8                   1000000000               0.3311 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8                    1000000000               0.3271 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8                    1000000000               0.3350 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8                    1000000000               0.3579 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8                    1000000000               0.3266 ns/op          0 B/op          0 allocs/op
BenchmarkResult_Ok-8                    1000000000               0.3218 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8                 1000000000               0.3346 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8                 1000000000               0.3384 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8                 1000000000               0.3196 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8                 1000000000               0.3072 ns/op          0 B/op          0 allocs/op
BenchmarkResultTestOk-8                 1000000000               0.3384 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyErr-8             1000000000               0.3332 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyErr-8             1000000000               0.3441 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyErr-8             1000000000               0.3591 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyErr-8             1000000000               0.3356 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyErr-8             1000000000               0.3372 ns/op          0 B/op          0 allocs/op
BenchmarkAsResultEmptyNoErr-8            9647712               126.9 ns/op            32 B/op          1 allocs/op
BenchmarkAsResultEmptyNoErr-8            9475947               122.0 ns/op            32 B/op          1 allocs/op
BenchmarkAsResultEmptyNoErr-8           10042959               121.4 ns/op            32 B/op          1 allocs/op
BenchmarkAsResultEmptyNoErr-8            8818743               128.3 ns/op            32 B/op          1 allocs/op
BenchmarkAsResultEmptyNoErr-8            9448839               123.7 ns/op            32 B/op          1 allocs/op
BenchmarkResult_OkInt-8                 75673644                15.90 ns/op            0 B/op          0 allocs/op
BenchmarkResult_OkInt-8                 67860129                15.87 ns/op            0 B/op          0 allocs/op
BenchmarkResult_OkInt-8                 72797239                15.68 ns/op            0 B/op          0 allocs/op
BenchmarkResult_OkInt-8                 67013431                15.60 ns/op            0 B/op          0 allocs/op
BenchmarkResult_OkInt-8                 75403287                15.89 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyInt-8              77981290                15.62 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyInt-8              65899971                15.50 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyInt-8              65190609                15.89 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyInt-8              73411852                16.75 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyInt-8              74789524                15.55 ns/op            0 B/op          0 allocs/op
BenchmarkResult_OkString-8              14390919                80.40 ns/op           16 B/op          1 allocs/op
BenchmarkResult_OkString-8              13817521                76.45 ns/op           16 B/op          1 allocs/op
BenchmarkResult_OkString-8              14965113                75.68 ns/op           16 B/op          1 allocs/op
BenchmarkResult_OkString-8              14652274                76.67 ns/op           16 B/op          1 allocs/op
BenchmarkResult_OkString-8              15157803                74.35 ns/op           16 B/op          1 allocs/op
BenchmarkResult_EmptyString-8           64723770                16.29 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyString-8           75101718                16.66 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyString-8           75580552                16.07 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyString-8           73452633                17.00 ns/op            0 B/op          0 allocs/op
BenchmarkResult_EmptyString-8           66919212                16.18 ns/op            0 B/op          0 allocs/op
BenchmarkOption_Some-8                  1000000000               0.3170 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8                  1000000000               0.3457 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8                  1000000000               0.3258 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8                  1000000000               0.3376 ns/op          0 B/op          0 allocs/op
BenchmarkOption_Some-8                  1000000000               0.3390 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8                  1000000000               0.3127 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8                  1000000000               0.3461 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8                  1000000000               0.3179 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8                  1000000000               0.3249 ns/op          0 B/op          0 allocs/op
BenchmarkOption_None-8                  1000000000               0.3378 ns/op          0 B/op          0 allocs/op
BenchmarkOption_IsNone-8                 9615813               120.5 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsNone-8                 9176836               126.9 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsNone-8                 9804849               127.3 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsNone-8                 9750918               121.7 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsNone-8                 9103474               122.2 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsSome-8                10008888               122.0 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsSome-8                 9285494               124.5 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsSome-8                10085062               123.7 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsSome-8                 9960693               126.5 ns/op            32 B/op          1 allocs/op
BenchmarkOption_IsSome-8                 9039792               128.1 ns/op            32 B/op          1 allocs/op
BenchmarkErrorCheck-8                   1000000000               0.3217 ns/op          0 B/op          0 allocs/op
BenchmarkErrorCheck-8                   1000000000               0.3428 ns/op          0 B/op          0 allocs/op
BenchmarkErrorCheck-8                   1000000000               0.3581 ns/op          0 B/op          0 allocs/op
BenchmarkErrorCheck-8                   1000000000               0.3357 ns/op          0 B/op          0 allocs/op
BenchmarkErrorCheck-8                   1000000000               0.3328 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/JPratama7/safe       84.108s
```

### Note
Error and None methods usable as structless but it doesn't infere types so instead of using `safetypes.None[T]()` and `safetypes.Err[T]("")` you could use them as how in examples above
