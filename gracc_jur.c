#define Py_LIMITED_API
#include <Python.h>

PyObject * convert(PyObject *, PyObject *);

// Workaround missing variadic function support
// https://github.com/golang/go/issues/975
int PyArg_ParseTuple_S(PyObject * args, char ** a) {  
    return PyArg_ParseTuple(args, "s", a);
}

static PyMethodDef GraccMethods[] = {  
    {"convert", convert, METH_VARARGS, "Convert JUR to json"},
    {NULL, NULL, 0, NULL}
};

PyMODINIT_FUNC  
initgracc(void)  
{
    (void) Py_InitModule("gracc", GraccMethods);
    //return PyModule_Create(&foomodule);
}
