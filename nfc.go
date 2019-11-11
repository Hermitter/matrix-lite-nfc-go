package nfc

/*
#cgo CXXFLAGS: -std=c++11 -DNXPBUILD__PH_RASPBERRY_PI -O3 -I/usr/local/include/matrix_nfc/nxp_nfc/NxpNfcRdLib/types -I/usr/local/include/matrix_nfc/nxp_nfc/NxpNfcRdLib/intfs
#cgo LDFLAGS:  -lmatrix_hal_nfc

#include <stdlib.h>
#include "nfc.h"
*/
import (
	"C"
)
import "unsafe"

func Status(code int) string {
	status := C.nfc_status(C.int(code))
	defer C.free(unsafe.Pointer(status))

	return C.GoString(status)
}
