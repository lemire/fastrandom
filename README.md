## fastrandom

Fast random number generation in Go. We use a division-less range function
to get a "fast" shuffle in Go.



```bash
$ go get github.com/davidminor/gorand/pcg
$ go test -bench=.
testing: warning: no tests to run
BenchmarkStandardShuffleWithGo1000-2                        50000         30891 ns/op
BenchmarkStandardShuffleWithPCGWithDivision1000-2          100000         12006 ns/op
BenchmarkStandardShuffleWithPCGButNoDivision1000-2         200000          8803 ns/op
```

Further reading: 
http://lemire.me/blog/2016/06/30/fast-random-shuffling/
