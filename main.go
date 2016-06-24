package main

import (
	"github.com/opensciencegrid/gracc-collector/gracc"
)

// #cgo CFLAGS: -I/usr/include/python2.6 -I/usr/include/python2.6 -fno-strict-aliasing -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -m64 -mtune=generic -D_GNU_SOURCE -fPIC -fwrapv -DNDEBUG -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -m64 -mtune=generic -D_GNU_SOURCE -fPIC -fwrapv
// #cgo LDFLAGS: -lpthread -ldl -lutil -lm -lpython2.6
// #include <Python.h>
// int PyArg_ParseTuple_S(PyObject *, char **);
import "C"


//export convert
func convert(self, args *C.PyObject) *C.PyObject {  
    var jur *C.char

    // Make sure the arguments are correct
    if C.PyArg_ParseTuple_S(args, &jur) == 0 {
        return nil
    }
    // Convert the JUR to Go string
    var go_jur string
    go_jur = C.GoString(jur)
    // Convert the Go string to byte array
    byteArray := []byte(go_jur)

    // Parse the byte array with the JobUsageRecord
    var v gracc.JobUsageRecord
    v.ParseXML(byteArray)
    j, _ := v.ToJSON("    ")

    // Convert the byte array back to string and return
    s := string(j[:])
    return C.PyString_FromString(C.CString(s))
}

func main() {}  
