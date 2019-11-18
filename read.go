package nfc

/*
#include <stdlib.h>
#include "./hal_nfc_wrapper/read.h"
*/
import (
	"C"
)

import (
	"unsafe"
	"fmt"
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

	// Convert Go struct to C
	var cOptions C.nfc_read_conf
	cOptions.info = C.bool(options.Info)
	cOptions.pages = C.bool(options.Pages)
	cOptions.ndef = C.bool(options.Ndef)
	cOptions.page = C.int(0)

	// Grab read results
	var cResults C.nfc_read_result = C.nfc_read(cOptions)

	// Convert C results to Go
	results := ReadResult{}
	if options.Info {
		populateInfo(&results, cResults.info)
	}
	if options.Page > 0 {
		// TODO fix
		// pointer := unsafe.Pointer(cResults.page)
		array := C.GoBytes(unsafe.Pointer(cResults.page), C.int(4))
		fmt.Println("\n", array)
		// C.free(pointer)

		// populatePage(&results, &cResults)
	}

	return int(cResults.status), results
}
