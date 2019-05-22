# gowasm
Sample WASM calculations in Go

## Usage
HTML and JS code can be found in the wasm_exec.html file. index.html is a symlink to wasm_exec.html.
Go code can be found in main.go. To compile to WASM:

```sh
GOOS=js GOARCH=wasm go build -o test.wasm
```

The test.wasm file is the actual binary to be run by the browser VM. You can use my simple web server at
https://github.com/jecolon/ws to run this sample. The following assumes your $GOPATH/bin directory is in your
$PATH.

```sh
git clone https://github.com/jecolon/gowasm
cd gowasm
go get github.com/jecolon/ws
ws
```
Then you can go to http://localhost:8080 in your browser to see the demo. Click the "Run" button to begin.

NOTE: wasm_exec.js and wasm_exec.html originals can be found in your Go installation directory under the misc/wasm subdirectory. Thay should be used with the Go version they came with.
