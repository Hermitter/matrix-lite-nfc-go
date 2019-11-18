package nfc

/*
#include <stdlib.h>
#include "./hal_nfc_wrapper/read.h"
*/
import (
	"C"
)
import (
	// "fmt"
	// "unsafe"
)

func populateInfo(r *ReadResult, info C.nfc_info_data) {
	r.Info = ReadInfoData{}
	r.Info.Technology = cGoString(info.technology)
	r.Info.Type = cGoString(info.card_type)
	r.Info.UID = cGoString(info.UID)
	r.Info.ATQ = cGoString(info.ATQ)
	r.Info.SAK = cGoString(info.SAK)
	r.Info.CardFamily = cGoString(info.card_family)
	r.Info.ICType = cGoString(info.IC_type)
	r.Info.BitRate = int(info.bit_rate)
	r.Info.StorageSize = int(info.storage_size)
	r.Info.ReadStatus = int(info.read_status)
}

// TODO fix
func populatePage(r *ReadResult, res *C.nfc_read_result) {	
	// pointer := unsafe.Pointer(res.page)
	// array := C.GoBytes(pointer, C.int(4))
	// fmt.Println(array)
	// C.free(pointer)

	r.Page = [4]uint{1, 2, 3, 4}
}
