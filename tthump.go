package main

// #cgo pkg-config: tthump
// #include <tthump.h>
import "C"

import "fmt"

func main() {
	fmt.Println("vim-go")

	var hndl: *C.tth

	C.tth_create(&hndl, 4)
}
