#include "./nfc.hpp"
#include <iostream>
#include <string.h>

extern "C" {
    #include "nfc.h"
}

matrix_hal::NFC nfc;
matrix_hal::NFCData nfc_data;

namespace helper {
    // - Convert a string to char *
    char* str_to_char(std::string str) {
        return strdup(&*str.begin());
    }
}

// Returns a string for the given status code.
char* nfc_status(int status_code) {
    return helper::str_to_char(matrix_hal::NFCStatus(status_code));
}