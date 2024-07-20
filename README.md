go-unmarshalperformance-78772815
================================

This repo contains the code for [my answer](https://stackoverflow.com/a/78772977/1296707)
to the question ["Optimizing json.Marshal in Go"](https://stackoverflow.com/questions/78772815/optimizing-json-marshal-in-go) on [StackOverflow](https://stackoverflow.com).

Usage
-----

Simply check out this repository and run `go test -bench=.` from it's root.

If you want to see how the behavior changes with various CPU sizes, you can use `go test -bench=. -cpu=1,2,4` to test for 1, 2 and 4 CPUs. You can use as many CPU Cores as your physical system has.

### Example output

    go test -bench=. -cpu=1,2,4
    goos: darwin
    goarch: arm64
    pkg: github.com/mwmahlberg/go-unmarshalperformance-78772815
    Benchmark64kMapStringInterface              5490            217707 ns/op
    Benchmark64kMapStringInterface-2            6572            172025 ns/op
    Benchmark64kMapStringInterface-4            6513            176681 ns/op
    Benchmark128kMapStringInterface             1321            893114 ns/op
    Benchmark128kMapStringInterface-2           1632            707845 ns/op
    Benchmark128kMapStringInterface-4           1604            737196 ns/op
    Benchmark256kMapStringInterface             1322            901109 ns/op
    Benchmark256kMapStringInterface-2           1635            706265 ns/op
    Benchmark256kMapStringInterface-4           1606            730651 ns/op
    Benchmark64kStruct                         16815             72051 ns/op
    Benchmark64kStruct-2                       16618             69919 ns/op
    Benchmark64kStruct-4                       16864             70883 ns/op
    Benchmark128kStruct                         3979            289447 ns/op
    Benchmark128kStruct-2                       4040            280413 ns/op
    Benchmark128kStruct-4                       4089            286967 ns/op
    Benchmark256kStruct                         4029            283122 ns/op
    Benchmark256kStruct-2                       4070            282152 ns/op
    Benchmark256kStruct-4                       4094            285352 ns/op
    Benchmark64kEncoder                         6207            212767 ns/op
    Benchmark64kEncoder-2                       6490            172975 ns/op
    Benchmark64kEncoder-4                       6692            174391 ns/op
    Benchmark128kEncoder                        1358            867077 ns/op
    Benchmark128kEncoder-2                      1651            698896 ns/op
    Benchmark128kEncoder-4                      1638            716011 ns/op
    Benchmark256kEncoder                        1358            872206 ns/op
    Benchmark256kEncoder-2                      1656            706698 ns/op
    Benchmark256kEncoder-4                      1647            721699 ns/op
    PASS
    ok      github.com/mwmahlberg/go-unmarshalperformance-78772815  36.178s