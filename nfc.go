package nfc

/*
#cgo CXXFLAGS: -std=c++11 -DNXPBUILD__PH_RASPBERRY_PI -O3 -I/usr/local/include/matrix_nfc/nxp_nfc/NxpNfcRdLib/types -I/usr/local/include/matrix_nfc/nxp_nfc/NxpNfcRdLib/intfs
#cgo LDFLAGS:  -lmatrix_hal_nfc

#include "nfc.h"
#include <stdlib.h>
*/
import (
	"C"
)
import "unsafe"

// Status returns a message from an NFC status code.
func Status(code int) string {
	return cGoString(C.nfc_status(C.int(code)))
}

// Handle CString to GoString conversion & memory deallocation
func cGoString(cstring *C.char) string {
	defer C.free(unsafe.Pointer(cstring))
	return C.GoString(cstring)
}

// TODO: add a mutex lock for NFC usage
// https://tour.golang.org/concurrency/9
