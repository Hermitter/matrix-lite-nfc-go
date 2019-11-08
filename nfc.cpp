#include "nfc.hpp"
#include <iostream>

extern "C" {
    #include "nfc.h"
}

namespace helper {
    // - Convert a string to char *
    char * str_to_char(std::string str) {
        return &*str.begin();
    }
}

char * nfc_status(int status_code) {
    return helper::str_to_char(matrix_hal::NFCStatus(status_code));
}