#### memory profile

Profiling a WebAssembly (Wasm) application using wzprof is a straightforward process. wzprof is a profiler based on the Wazero runtime, designed to collect CPU and memory profiles for Wasm applications.

build and compile the go code to wasm file using wasip1 port of go

```shell
GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go
```

we can build it using js port too

```shell
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

```shell
wzprof -sample 1 -memprofile /Users/rajiv.singh/Desktop/hello-wasm/memprofile.out ./main.wasm
```

`/Users/rajiv.singh/Desktop/hello-wasm/memprofile.out` is the path where memory profile will be saved.

#### cpu profile

```shell
wzprof -sample 1 -cpuprofile /Users/rajiv.singh/Desktop/hello-wasm/cpuprofile ./main.wasm
```

also, pprof profiles the whole application so we can see lot of profile when we rn above things we still need to initialize the go runtime to get the profile. The Go runtime is just more Go code that runs in the same process as your code. It's not a separate process that can be profiled separately. So, we need to initialize the Go runtime to get the profile.

---

why i am using wazero to build the wasm file?
By avoiding CGO, wazero avoids prerequisites such as shared libraries or libc, and lets you keep features like cross compilation. Being pure Go, wazero adds only a small amount of size to your binary. Meanwhile, wazero’s API gives features you expect in Go, such as safe concurrency and context propagation.

---

wazero:

- When a runtime is written in Go, it means that the code for the runtime is written using the Go programming language. This allows the runtime to be compiled into a Go library. so wazero is a runtime written in Go that can be compiled into a Go library

---

Because we use expose the profiles on pprof http endpoints, _any continuous-profiler system supporting pprof profiles is compatible with wzprof_.


---

wasmimport directive: imports function to WebAssembly module

wasmexport directive: Export function to wasm module 

Like a go function to calculate something can be called from wasm

