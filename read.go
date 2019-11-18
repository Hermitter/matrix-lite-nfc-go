package nfc

/*
#include "hal_nfc_wrapper/read.h"
*/
import (
	"C"
)

// ReadConf determines what data is read from a tag
type ReadConf struct {
	Info, Pages, Ndef bool
	Page              int
}

type ReadInfoData struct {
	Technology, Type, UID, ATQ, SAK, CardFamily, ICType string
	BitRate, StorageSize, ReadStatus                    int
}

type ReadResult struct {
	Info ReadInfoData
	Page [4]uint
	// pages ReadPagesData
	// ndef	 ReadNdefData
}

// Read scans an NFC tag
func Read(options ReadConf) (int, ReadResult) {
	// lock/unlock mutex
	if !nfcLock.TryLock() {
		return 1024, ReadResult{}
	}
	// nfcLock.lock()
	defer nfcLock.Unlock()

	// Convert Go struct to C struct
	var cOptions C.nfc_read_conf
	cOptions.info = C.bool(options.Info)
	cOptions.pages = C.bool(options.Pages)
	cOptions.ndef = C.bool(options.Ndef)
	cOptions.page = C.int(0)

	// Grab read results
	var cResults C.nfc_read_result = C.nfc_read(cOptions)

	// Convert C struct to Go struct
	results := ReadResult{}
	populateInfo(&results, cResults.info)
	results.Page = [4]uint{2, 3, 5, 4}

	return int(cResults.status), results
}

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
