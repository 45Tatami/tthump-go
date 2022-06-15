package main

/*
#cgo pkg-config: tthump
#include <tthump.h>
#include <stdlib.h>
#include <stdint.h>

void cgo_callback_glue(int64_t job, char *thmbFile, void *uparam);
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

type TthRequest struct {
	ImgPath string
}

type TthResult struct {
	Ok       bool
	ThmbPath string
	Request  *TthRequest
}

type callbackUparam struct {
	req  *TthRequest
	cres chan *TthResult
}

//export goCallback
func goCallback(job int64, thmbPathC *C.char, handle C.uintptr_t) {
	thmbPathGo := C.GoString(thmbPathC)
	handleGo := cgo.Handle(handle)
	uparam := handleGo.Value().(callbackUparam)

	uparam.cres <- &TthResult{true, thmbPathGo, uparam.req}
	handleGo.Delete()
}

func Create(n_workers uint8) (chan<- *TthRequest, <-chan *TthResult) {
	creq := make(chan *TthRequest)
	cres := make(chan *TthResult)

	go func() {
		var hndl *C.struct_tth
		C.tth_create(&hndl, C.uint8_t(n_workers))
		for {
			req, ok := <-creq
			if !ok || req == nil {
				break
			}

			uparamC := cgo.NewHandle(callbackUparam{req, cres})

			cpath := C.CString(req.ImgPath)

			C.tth_get_thumbnail_async(hndl, cpath,
				(*C.tth_callback)(
					unsafe.Pointer(C.cgo_callback_glue)),
				unsafe.Pointer(uparamC))

			C.free(unsafe.Pointer(cpath))
		}
		C.tth_destroy(hndl)
	}()

	return creq, cres
}
