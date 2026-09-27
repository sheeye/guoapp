package main

/*
#include "crash.h"
*/
import "C"

//export DuanjuSetCrashDir
func DuanjuSetCrashDir(dir *C.char) {
	C.duanju_set_crash_dir(dir)
}

func init() {
	C.init_crash_dir_from_env()
}
