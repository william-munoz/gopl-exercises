# ch03/ex08

## Performance and memory usage

Performance and memory usage are as follows, with complex64 as 1.
The smaller the better.

```txt
            Performance     Memory
complex64          1.00       1.00
complex128         0.89       1.00
big.Float         95.22   33040.00
big.Rat          168.39   55872.03
```

```txt
$ go test -benchmem -bench=.
BenchmarkMandelbrotComplex64-4            100000             18338 ns/op              30 B/op         25 allocs/op
BenchmarkMandelbrotComplex128-4           100000             16400 ns/op              30 B/op         25 allocs/op
BenchmarkMandelbrotBigFloat-4               1000           1746090 ns/op          991200 B/op      11827 allocs/op
BenchmarkMandelbrotBigRat-4                  500           3087937 ns/op         1676161 B/op      36621 allocs/op
PASS
ok      github.com/kdama/gopl/ch03/ex08 7.639s
```

## Rendering result

At 2 ^ 15x magnification, no difference was seen in the rendering results of complex64 and complex128. (`examples/complex64_z15.png`, `examples/complex128_z15.png`)
At 2 ^ 16x magnification, we could see the difference in the rendering results of complex64 and complex128. (`examples/complex64_z16.png`, `examples/complex128_z16.png`)

At 2 ^ 44x magnification, no difference was seen in the rendering results of complex128 and big.Float. (`examples/complex128_z44.png`, `examples/bigfloat_z44.png`)
At 2 ^ 45x magnification, we could see a difference in the rendering results of complex128 and big.Float. (`examples/complex128_z45.png`, `examples/bigfloat_z45.png`)

The difference between the rendering results of big.Float and big.Rat could not be obtained in detail.
At 2 ^ 52x magnification, no difference was found in the rendering results of big.Float and big.Rat. (`examples/bigfloat_z52.png`, `examples/bigrat_z52.png`)
At 2 ^ 53x magnification, we could see the difference in the rendering results of big.Float and big.Rat. (`examples/bigfloat_z53.png`, `examples/bigrat_z53.png`)
However,
At 2 ^ 53 times magnification, both big.Float and big.Rat gave clearly wrong rendering results that seemed to exceed the accuracy limit.
