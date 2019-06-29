

#include "go-py.h"

void PyDecRef(RefPyObject obj) {

    if (obj != NULL) {
        Py_DECREF(obj);
        obj = NULL;
        return ;
    }
}

void _PyDecRef(RefPyObject *obj) {

    if (*obj != NULL) {
        Py_DECREF(*obj);
        *obj = NULL;
        return ;
    }
}

void PyExec(const char *cmdString) {
    PyRun_SimpleString(cmdString);
    return;
}

PyModule PyRequireModule(const char *moduleName) {
    PyModule mModule = PyImport_ImportModule(moduleName);
    if(!mModule){
        printf("import python module failed!!\n");
        PyErr_Print();
        exit(ImportModuleError);
    }
    return mModule;
}

RefPyObject PyGetAttr(PyModule mModule, const char *objectName) {
    RefPyObject mObject = PyObject_GetAttrString(mModule, objectName);

    if(!mObject){
        printf("import python Object failed!!\n");
        PyErr_Print();
        exit(ImportObjectError);
    }
    return mObject;
}

RefPyObject PyRequireObject(const char *moduleName, const char *objectName) {
    PyModule mModule = PyRequireModule(moduleName);
    if(!mModule){
        printf("import python module failed!!\n");
        PyErr_Print();
        exit(ImportObjectError);
    }

    RefPyObject mObject = PyGetAttr(mModule, objectName);
    _PyDecRef(&mModule);
    return mObject;
}

RefPyObject PyInvokeMemberFunction(PyModule mModule, const char *funcName, PyTuple pArgs) {
    PyFunction pFunc = PyGetAttr(mModule, funcName);
    RefPyObject pRet = PyObject_CallObject(pFunc, pArgs);
    _PyDecRef(&pFunc);
    return pRet;
}

const char *PyObjectAsString(RefPyObject pObj) {
    return _PyUnicode_AsString(pObj);
}

int PyIsString(RefPyObject pObj) {
    return PyUnicode_Check(pObj);
}

void PyPrintObject(RefPyObject pObj) {
    PyModule mainModule = PyEval_GetGlobals();
    PyTuple pArgs = PyTuple_New(1);
    PyTuple_SetItem(pArgs, 0, pObj);

    if(!mainModule)    //检查函数的有效性
    {
        printf("Can't require module!\n");
        PyErr_Print();
        return ;
    }

    PyFunction pFunc = PyDict_GetItemString(mainModule, "print");
    
    if(!pFunc || !PyCallable_Check(pFunc))    //检查函数的有效性
    {
        printf("Can't find function!\n");
        PyErr_Print();
        return ;
    }
    PyExec("print(\"Here\")");
    PyObject_CallObject(pFunc, pArgs);
    _PyDecRef(&pFunc);
    _PyDecRef(&pArgs);
    return ;
}

RefPyObject PyGetItem(PyTuple pTuple, long long t) {
    return PyTuple_GetItem(pTuple, t);
}



void PyTimeIt() {
    PyRun_SimpleString("from time import time,ctime\n"
		               "print('Today is', ctime(time()))\n");
    return ;
}

void PyInit() {
    Py_Initialize();
	if (!Py_IsInitialized())
	{

	    printf("Python init failed!\n");
        PyErr_Print();
	    exit(-1);
	}
}

PyString NewPyString(const char *goStr) {
    return Py_BuildValue("s", goStr);
}

