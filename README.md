# cabinet
A distribute memory cache storage for Go

## Benchmark
### HTTP

#### Set
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
#### GET
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

### TCP

#### Set
```text
% ./benchmark -type tcp -n 100000 -r 100000 -t set
type is tcp
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
3.085778 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 6 ms
99% requests < 9 ms
99% requests < 24 ms
99% requests < 28 ms
99% requests < 29 ms
100% requests < 128 ms
29 usec average for each request
throughput is 32.406734 MB/s
rps is 32406.734330

```

#### GET

```text
% ./benchmark -type tcp -n 100000 -r 100000 -t get
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 1
63059 records get
36941 records miss
0 records set
2.951104 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 25 ms
99% requests < 28 ms
100% requests < 42 ms
28 usec average for each request
throughput is 21.367935 MB/s
rps is 33885.622385

```