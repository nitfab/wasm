set GOOS=js
set GOARCH=wasm
cd cmd
cd wasm
go build -o ../../assets/main.wasm  <!-- It will create a main.wasm file in assets folder -->
It will also create a js file where we have go root.
You can see your GOROOT location by cmd: go env GOROOT
TO copy that file inside your assets folder. Just cd to assets folder and write this cmd: 
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" <!-- It will copy the wasm_exec.js file from the path-->
Now include the exce.js file into the script tag in index.html: <script src="wasm_exec.js"></script> 

https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface/instantiateStreaming

Run the file: cmd\server> go run main.go with that run index.html(open with live server)

