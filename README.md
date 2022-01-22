# go fmt in the browser via WebAssembly

**DEMO: [go-fmt-wasm.netlify.app](https://go-fmt-wasm.netlify.app)**

This example uses the `go/format` package instead of compiling the `go fmt` command line tool because we don't have access to a file-system replacement out of the box in the browser which `go fmt` would need.

[TinyGo](https://tinygo.org/) produces significantly smaller wasm files so we use that instead of the the standard Go compiler.  
Compared to the standard Go version, the size of the wasm file drops from ~3MB to ~0,5MB (uncompressed).
The time from initial page load to when the format button can be pressed was reduced from ~1s to ~0.5s.

## Creating the wasm file
- Install tinygo, see https://tinygo.org/getting-started/install/
- There is a custom wasm compile configuration needed (`wasm.json`) otherwise the stack size limit is too low to perform the formatting when there are more then ~100 (?) lines of code.
  With the current setting in the config, 3000 lines of code can be formatted without problems (takes ~130ms).
- Compile the Go code with `tinygo build -o format-go-code.wasm -target ./wasm.json -no-debug ./go/main.go`
- To execute the wasm file from JS you also need some glue code. This code can also be found in your local go installation under `{path to your tinygo installation}/targets/wasm_exec.js`. This script needs to be imported as a script in your index.html. (On Linux, the Go directory is usually `/usr/local/lib/tinygo`.)
- To avoid a memory leak, one function in the `wasm_exec.js` file needs to be replaced, see https://github.com/tinygo-org/tinygo/issues/1140#issuecomment-718145455

## Using the wasm file
- The glue code needs to be imported via `<script src="wasm_exec.js"></script>`.
- The WASM file needs to be loaded, see index.html.
- Afterwards, the "exported" function (`formatGoCode`) is available in JavaScript.

## Error handling
- If the formatting did not succeed, an empty string is returned. In that case, the result should not be used to replace the non-formatted code the user wrote.

## Sources
- https://tinygo.org/docs/guides/webassembly/
- https://github.com/tinygo-org/tinygo/tree/release/src/examples/wasm/slices

## Motivation
This is a standalone prototype for https://github.com/exercism/exercism/issues/5986.