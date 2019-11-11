package nfc

/*
#include "read.h"
*/
import (
	"C"
)

// ReadConf determines what data is read from a tag
type ReadConf struct {
	Info, Pages, Ndef bool
	Page              uint
}

type ReadInfoData struct {
	Technology, Type, UID, ATQ, SAK, CardFamily, ICType string
	BitRate, StorageSize, ReadStatus                    int
}

type ReadResult struct {
	Info ReadInfoData
}

// Read scans an NFC tag
func Read(options ReadConf) (int, ReadResult) {
	// Convert Go struct to C struct
	var cOptions C.nfc_read_conf
	cOptions.info = C.bool(options.Info)
	cOptions.pages = C.bool(options.Pages)
	cOptions.ndef = C.bool(options.Ndef)
	cOptions.page = C.int(0)

	// Convert C struct to Go struct
	var cResults C.nfc_read_result = C.nfc_read(cOptions)

	results := ReadResult{}
	populateInfo(&results, cResults.info)

	return int(cResults.status), results
}

func populateInfo(r *ReadResult, info C.nfc_info_data) {
	r.Info = ReadInfoData{}
	r.Info.Technology = C.GoString(info.technology)
	r.Info.Type = C.GoString(info.card_type)
	r.Info.UID = C.GoString(info.UID)
	r.Info.ATQ = C.GoString(info.ATQ)
	r.Info.SAK = C.GoString(info.SAK)
	r.Info.CardFamily = C.GoString(info.card_family)
	r.Info.ICType = C.GoString(info.IC_type)
	r.Info.BitRate = int(info.bit_rate)
	r.Info.StorageSize = int(info.storage_size)
	r.Info.ReadStatus = int(info.read_status)
}
