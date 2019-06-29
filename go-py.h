
#ifndef MYD_GOPY
#define MYD_GOPY

#include "pylib3.7/include/Python.h"
#pragma comment(lib, "./pylib3.7/libs/libpython37.a")

#define ImportModuleError 1001
#define ImportObjectError 1002
#define ExecFunctionError 1003


typedef PyObject *RefPyObject;
typedef RefPyObject PyFunction;
typedef RefPyObject PyModule;
typedef RefPyObject PyTuple;
typedef RefPyObject PyDict;
typedef RefPyObject PyError;
typedef RefPyObject PyString;
typedef RefPyObject PyClass;

void PyDecRef(RefPyObject obj);

void PyExec(const char *cmdString);

PyModule PyRequireModule(const char *moduleName);

RefPyObject PyGetAttr(PyModule mModule, const char *objectName);

RefPyObject PyRequireObject(const char *moduleName, const char *objectName);

RefPyObject PyInvokeMemberFunction(PyModule mModule, const char *funcName, PyTuple pArgs);

const char *PyObjectAsString(RefPyObject pObj);

int PyIsString(RefPyObject pObj);

void PyPrintObject(RefPyObject pObj);

RefPyObject PyGetItem(PyTuple pTuple, long long t);


void PyTimeIt();

void PyInit();

PyString NewPyString(const char *goStr);


#endif
