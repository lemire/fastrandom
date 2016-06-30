## fastrandom

Fast random number generation in Go. We use a division-less range function
to get a "fast" shuffle in Go.



```bash
$ go get github.com/dgryski/go-pcgr
$ go version
go version go1.7beta2 linux/amd64
$ go test -bench=.
testing: warning: no tests to run
BenchmarkStandardShuffleWithGo1000-2                                50000         30628 ns/op
BenchmarkStandardShuffleWithPCGWithDivision1000_dgryski-2          200000          9016 ns/op
BenchmarkStandardShuffleWithPCGButNoDivision1000_dgryski-2         300000          5066 ns/op
PASS
ok      _/home/dlemire/CVS/github/fastrandom    5.293s
```

Further reading: 
http://lemire.me/blog/2016/06/30/fast-random-shuffling/
