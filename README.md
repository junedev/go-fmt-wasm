# go fmt in the browser via WebAssembly

**DEMO: [go-fmt-wasm.netlify.app](go-fmt-wasm.netlify.app)**

This example uses the `go/format` package instead of compiling the `go fmt` command line because we don't have access to a file-system replacement out of the box in the browser which `go fmt` would need.

[Tinygo](https://tinygo.org/) produces significantly smaller wasm files so we use that instead of the the standard Go compiler.

## Creating the wasm file
- Install tinygo, see https://tinygo.org/getting-started/install/
- Compile the Go code with `tinygo build -o format-go-code.wasm -target wasm -no-debug ./go/main.go`
- To execute the wasm file from JS you also need some glue code. This code can also be found in your local go installation under `{path to your tinygo installation}/targets/wasm_exec.js`. This script needs to be imported as a script in your index.html. (On Linux, the Go directory is usually `/usr/local/lib/tinygo`.)
- To avoid a memory leak, one function in the `wasm_exec.js` file needs to be replaced, see https://github.com/tinygo-org/tinygo/issues/1140#issuecomment-718145455

## Using the wasm file
- The glue code needs to be imported via `<script src="wasm_exec.js"></script>`.
- The WASM file needs to be loaded, see index.html.
- Afterwards, the "exported" function (`formatGoCode`) is available in JavaScript.

## Sources
- https://tinygo.org/docs/guides/webassembly/
- https://github.com/tinygo-org/tinygo/tree/release/src/examples/wasm/slices
