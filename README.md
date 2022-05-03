# cabinet
A distribute memory cache storage for Go

## Benchmark

```text
% ./cache-benchmark -type http -n 100000 -r 100000 -t set
type is http
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is set
keyspacelen is 100000
pipeline length is 1
0 records get
0 records miss
100000 records set
3.830248 seconds total
99% requests < 1 ms
100% requests < 7 ms
37 usec average for each request
throughput is 26.107967 MB/s
rps is 26107.966910
```

```text
% ./cache-benchmark -type http -n 100000 -r 100000 -t get
type is http
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 1
62926 records get
37074 records miss
0 records set
3.723067 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
100% requests < 7 ms
36 usec average for each request
throughput is 16.901657 MB/s
rps is 26859.576918

```