#! /bin/bash

echo "========== <Running wasm script> =========="
GOOS=js GOARCH=wasm go build -o main.wasm

echo "======== copying wasm_exec_js file ======="
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

echo "======= creating html index file ========="
# echo "<html>
#     <head>
#         <meta charset="utf-8"/>
#         <script src="wasm_exec.js"></script>
#         <script>
#             const go = new Go();
#             WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
#                 go.run(result.instance);
#             });
#         </script>
#     </head>
#     <body></body>
# </html>" > index.html

echo "=============== <Done> ===================="