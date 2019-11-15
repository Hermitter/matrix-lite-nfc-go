#include "./nfc.hpp"

extern "C"
{
#include "./read.h"
}
#include <iostream>
#include <string.h>

nfc_read_result nfc_read(nfc_read_conf options)
{
    // Status codes from NFC reads
    int32_t nfc_status, info_status, pages_status, ndef_status;

    // Data to return from read
    nfc_info_data info_result{};
    uint8_t *nfc_page_data = nullptr;
    // nfc_pages_data pages_result{};
    // nfc_ndef_data ndef_result{};

    // Read requested data from NFC tag //
    if (options.info)
    {
        nfc_status = nfc.Activate();
        info_status = nfc.ReadInfo(&nfc_data.info);
        nfc.Deactivate();
    }
    if (options.page > -1)
    {
        nfc_status = nfc.Activate();

        // valid C array by passing pointer to start of vector
        nfc_page_data = nfc.mful.ReadPage(options.page).data();

        nfc.Deactivate();
    }

    // Populate remaining structs //
    if (nfc_data.info.recently_updated)
    {
        info_result.technology = helper::str_to_char(nfc_data.info.technology);
        info_result.card_type = helper::str_to_char(nfc_data.info.type);
        info_result.UID = helper::str_to_char(nfc_data.info.UIDToHex().erase(0, 2));
        info_result.ATQ = helper::str_to_char(nfc_data.info.ATQToHex().erase(0, 2));
        info_result.SAK = helper::str_to_char(nfc_data.info.SAKToHex().erase(0, 2));
        info_result.card_family = helper::str_to_char(nfc_data.info.card_family);
        info_result.IC_type = helper::str_to_char(nfc_data.info.IC_type);
        info_result.bit_rate = nfc_data.info.bit_rate;
        info_result.storage_size = nfc_data.info.storage_size;
        info_result.read_status = info_status;
    }

    return nfc_read_result{
        .info = info_result,
        .page = nfc_page_data,
        .status = nfc_status};
}