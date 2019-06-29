package GoPy


/*
#include "go-py.h"
*/
import "C"


func GetItem(tup PyTuple, t int64) RefPyObject {
	return C.PyGetItem(tup, C.longlong(t))
}
