## Golang and Web Assembly

A simple test of running go code in the web browser using web assembly

- The wasm code can be moved to a webworker to avoid blocking the main thread, especially,
when using the goroutines (I had to sleep them based on jittering the sleep seconds number randomly).

- On existing the main gorountine ith other gorountine in execution throws a non-terminal error which is expected of golang concurrency as other goroutines are in execution as the main gets terminated, forcing them to close unconditionally.https://permify.co/post/wasm-go/
