go-unmarshalperformance-78772815
================================

This repo contains the code for [my answer](https://stackoverflow.com/a/78772977/1296707)
to the question ["Optimizing json.Marshal in Go"](https://stackoverflow.com/questions/78772815/optimizing-json-marshal-in-go) on [StackOverflow](https://stackoverflow.com).

Usage
-----

Simply check out this repository and run `go test -bench=.` from it's root.

If you want to see how the behavir changes with various CPU sizes, you can use `go test -bench=. -cpu=1,2,4` to test for 1,2 and 4 CPUs.

### Example output

    $ go test -bench=. -cpu=1,2,4,6,8,10
    goos: darwin
    goarch: arm64
    pkg: github.com/mwmahlberg/go-unmarshalperformance-78772815
    Benchmark64k                2830            413607 ns/op
    Benchmark64k-2              2952            388968 ns/op
    Benchmark64k-4              2920            399129 ns/op
    Benchmark64k-6              2912            397847 ns/op
    Benchmark64k-8              2884            398933 ns/op
    Benchmark64k-10             2946            396289 ns/op
    Benchmark128k                663           1816750 ns/op
    Benchmark128k-2              704           1698653 ns/op
    Benchmark128k-4              696           1740807 ns/op
    Benchmark128k-6              626           1703486 ns/op
    Benchmark128k-8              699           1702414 ns/op
    Benchmark128k-10             698           1719027 ns/op
    Benchmark256k                652           1834965 ns/op
    Benchmark256k-2              699           1702141 ns/op
    Benchmark256k-4              690           1734279 ns/op
    Benchmark256k-6              682           1733939 ns/op
    Benchmark256k-8              700           1738449 ns/op
    Benchmark256k-10             692           1725643 ns/op
    PASS
    ok      github.com/mwmahlberg/go-unmarshalperformance-78772815  24.838s