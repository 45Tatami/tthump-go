package main

// #cgo pkg-config: tthump
// #include <tthump/tthump.h>
import "C"

import "fmt"

func main() {
	fmt.Println("vim-go")

	var hndl *C.struct_tth

	C.tth_create(&hndl, 4)

	j := C.tth_get_thumbnail_async(hndl, C.CString("/tmp/example.mkv"), nil, nil)
	C.tth_job_wait(hndl, j)
}
