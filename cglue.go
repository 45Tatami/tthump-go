package main

/*
#include <stdint.h>

void cgo_callback_glue(int64_t job, char *thmbFile, void *uparam) {
	void goCallback(int64_t, char *, void *);
	goCallback(job, thmbFile, uparam);
}
*/
import "C"
