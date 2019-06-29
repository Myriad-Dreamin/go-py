package GoPy


/*
#include "go-py.h"
*/
import "C"


func InvokeMemberFunction(mModule C.PyModule, funcName string, varList... interface{}) C.RefPyObject {
	// TODO: len is int -> int64 ? int32
	pArgs := C.RefPyObject(C.PyTuple_New(C.longlong(len(varList))))
	defer DecRef(&pArgs)
	for i := C.longlong(len(varList) - 1); i >= 0; i-- {
		C.PyTuple_SetItem(pArgs, i, varList[i].(RefPyObject))
	}
	return C.PyInvokeMemberFunction(mModule, C.CString(funcName), pArgs)
}

func AggregateObject(pObjArr... interface{}) []RefPyObject {
	var ret = make([]RefPyObject, 0, len(pObjArr))
	for _, pObj := range pObjArr {
		ret = append(ret, pObj.(RefPyObject))
	}
	return ret
}
