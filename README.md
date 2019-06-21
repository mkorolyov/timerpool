# timerpool

[![License MIT](https://img.shields.io/badge/License-MIT-blue.svg)](http://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/mkorolyov/timerpool?status.svg)](http://godoc.org/github.com/mkorolyov/timerpool) [![Go Report Card](https://goreportcard.com/badge/github.com/mkorolyov/timerpool)](https://goreportcard.com/report/github.com/mkorolyov/timerpool) [![Build Status](https://travis-ci.org/mkorolyov/timerpool.svg?branch=master)](http://travis-ci.org/mkorolyov/timerpool) [![codecov](https://codecov.io/gh/mkorolyov/timerpool/branch/master/graph/badge.svg)](https://codecov.io/gh/mkorolyov/timerpool)

sync.Pool of reusable timer.Timer with constant timeout passed.

#### Get the package using:

```
$ go get -u -v github.com/mkorolyov/timerpool
```

#### Usage

Create a pool with 
```go
pool := New(time.Second)
```

then in a hot peace of code acquire the timer and don't forget to release it.
```go
timer := pool.Acquire()
// use it
pool.Release(timer)
```