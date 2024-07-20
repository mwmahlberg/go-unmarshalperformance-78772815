go-unmarshalperformance-78772815
================================

This repo contains the code for [my answer](https://stackoverflow.com/a/78772977/1296707)
to the question ["Optimizing json.Marshal in Go"](https://stackoverflow.com/questions/78772815/optimizing-json-marshal-in-go) on [StackOverflow](https://stackoverflow.com).

Usage
-----

Simply check out this repository and run `go test -bench=.` from it's root.

If you want to see how the behavir changes with various CPU sizes, you can use `go test -bench=. -cpu=1,2,4` to test for 1,2 and 4 CPUs.