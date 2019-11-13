#ifndef NFC_HPP
#define NFC_HPP

#include "matrix_nfc/nfc_data.h"
#include "matrix_nfc/nfc.h"

// Global NFC Vars //
extern matrix_hal::NFC nfc;
extern matrix_hal::NFCData nfc_data;

// Helpful functions
namespace helper {
    char * str_to_char(std::string);
}

#endif