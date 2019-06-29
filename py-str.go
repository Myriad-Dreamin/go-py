package GoPy


/*
#include "go-py.h"
*/
import "C"


func PyString(goStr string) C.PyString {
	return C.NewPyString(C.CString(goStr))
}

func GoString(pObj C.RefPyObject) string {
	return C.GoString(C.PyObjectAsString(pObj))
}
