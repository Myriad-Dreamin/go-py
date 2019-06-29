package GoPy

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L/pylib3.7/libs -lpython37
#include "go-py.h"
*/
import "C"
import "os"
// import "fmt"
// import "os/signal"
// import "syscall"

type RefPyObject = C.RefPyObject
type PyModule = C.PyModule
type PyDict = C.PyDict
type PyTuple = C.PyTuple
type PyClass = C.PyClass

func DecRef(pObj *C.RefPyObject) {
	C.PyDecRef(*pObj)
	*pObj = nil
}

func FinalizeEx() int32 {
	return int32(C.Py_FinalizeEx())
}

func Exec(cmdString string) {
	C.PyExec(C.CString(cmdString))
}

func RequireModule(moduleName string) C.PyModule {
	return C.PyRequireModule(C.CString(moduleName))
}

func RequireObject(moduleName, objectName string) C.RefPyObject {
	return C.PyRequireObject(C.CString(moduleName), C.CString(objectName))
}

func GetAttr(mObj C.RefPyObject, attrName string) C.RefPyObject {
	return C.PyGetAttr(mObj, C.CString(attrName))
}

func PrintObject(pObj C.RefPyObject) {
	C.PyPrintObject(pObj)
}


func InvokeMemberFunction(mModule C.PyModule, funcName string, varList... interface{}) C.RefPyObject {
	// TODO: len is int -> int64 ? int32
	pArgs := C.RefPyObject(C.PyTuple_New(C.longlong(len(varList))))
	defer DecRef(&pArgs)
	for i := C.longlong(len(varList) - 1); i >= 0; i-- {
		C.PyTuple_SetItem(pArgs, i, varList[i].(RefPyObject))
	}
	return C.PyInvokeMemberFunction(mModule, C.CString(funcName), pArgs)
}

func PyString(goStr string) C.PyString {
	return C.NewPyString(C.CString(goStr))
}

func GoString(pObj C.RefPyObject) string {
	return C.GoString(C.PyObjectAsString(pObj))
}

func AggregateObject(pObjArr... interface{}) []RefPyObject {
	var ret = make([]RefPyObject, 0, len(pObjArr))
	for _, pObj := range pObjArr {
		ret = append(ret, pObj.(RefPyObject))
	}
	return ret
}

func GetItem(tup PyTuple, t int64) RefPyObject {
	return C.PyGetItem(tup, C.longlong(t))
}

type AtExitFunc func()

var atExitFuncs []AtExitFunc

func RegisterAtExitFunc(fuc AtExitFunc) {
	atExitFuncs = append(atExitFuncs, fuc)
}


func AtExit() {

	for _, ef := range atExitFuncs {
		ef()
	}

	if (FinalizeEx() < 0) {	
		os.Exit(120)
	}
	// return func() {
	// 	ch := make(chan os.Signal, 1)
	// 	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 	fmt.Println("here")
	// 	s := <-ch
	// 	fmt.Println("Got signal:", s)

	// 	fmt.Println("get", s)
		
	// 	// select {
	// 	// case 
	// 	// default :	
	// 	// 	fmt.Println("here")
	// 	// }
	// }
}

func init() {
	C.PyInit()
	Exec("import sys")
	Exec("sys.path.append('./')")

}
