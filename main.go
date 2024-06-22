package main

import (
	"fmt"
	"math/rand"
	"os"
	"syscall/js"
	"time"
)

var done = make(chan bool)

func main() {

	// fmt.Println("Hello, WebAssembly!")
	// js.Global().Set("newCall", newCall())
	// js.Global().Set("goCall", goCall())
	// <-make(chan struct{})

	// The actual issue is that when I write to the channel, it unblocks the 
	// `<-ch` in the default section of the for-select loop and the loop continues
	// The `case <- ch:` block never gets executed as the `<-ch` does the blocking/unblocking
	// for {
	// 	select {
	// 	case <-done:
	// 		fmt.Println("Exiting go application")
	// 		return
	// 		// break
	// 		// os.Exit(0)
	// 	default:
	// 		fmt.Println("Hello, WebAssembly!")
	// 		js.Global().Set("newCall", newCall())
	// 		js.Global().Set("goCall", goCall())
	// 		js.Global().Set("stop", stop())
	// 		<-done
	// 	}
	// }
	// os.Exit(0)

	fmt.Println("Hello, WebAssembly!")
	js.Global().Set("stop", stop())
	js.Global().Set("newCall", newCall())
	js.Global().Set("goCall", goCall())
	// closing this without wait for the other gorountines throws
	// some kind of error on exiting - the error is not terminal
	<-done
	fmt.Println("Exiting go application")
	os.Exit(0)
	

}

func stop() js.Func {
	st := js.FuncOf(func(this js.Value, args []js.Value) any {
		done <- true
		return "Stopping..........."
	})
	return st
}

func newCall() js.Func {
	new_call := js.FuncOf(func(this js.Value, args []js.Value) any {
		return "Result of heavy math computation on an image: " + args[0].String()
	})
	return new_call
}

func goCall() js.Func {
	rout := js.FuncOf(func(this js.Value, args []js.Value) any {
		// routine(1)
		// routine(2)
		// routine(3)
		for i := 1; i < 20; i++ {
			routine(i)
			fmt.Println("Time stamp: ", i, "started")
			// time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		}
		return nil
	})

	return rout
}

func routine(start int) {
	go func(start int) {
		for {
			fmt.Println("Time stamp: ", start, " ---> ", time.Now().Second())
			// time.Sleep(time.Second * time.Duration(start * rand.Intn(5)))
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		}
	}(start)
}
