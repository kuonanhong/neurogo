Portfolio Website on the Go
===========================

I always learn a new language by re-creating my simple
static website in that language; this time, I'm learning Go.

Go has piqued my interest primarily because of its support by Google,
its benchmarks (http://www.techempower.com/benchmarks/), docker (docker.com)
was written in Go, and the utilization of goroutines
(https://gobyexample.com/goroutines).  Having worked with Python for awhile now,
I've grown to enjoy coroutines for parallelization and concurrency as an
alternative to callback hell (http://elm-lang.org/learn/Escape-from-Callback-Hell.elm).

Build
-----

Navigate to `src` directory

```
go install
```

Build application

```
go build application.go
```

Or use gin which will allow automatic building of web app

```
go get https://github.com/codegangsta/gin
```

Run
---

```
./application
```

Or if using gin

```
gin
```

Navigate to http://localhost:3000/