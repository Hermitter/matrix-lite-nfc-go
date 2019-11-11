#ifndef READ_H
#define READ_H

#include <stdbool.h>

// Read options
typedef struct {
    bool info, pages, ndef;
    int page;
} nfc_read_conf;

// All NFC data //
typedef struct {
    char* technology;
    char* card_type;
    char* UID;
    char* ATQ;
    char* SAK;
    char* card_family;
    char* IC_type;
    int bit_rate;
    int storage_size;
    int read_status;
} nfc_info_data;
typedef struct {} nfc_pages_data;
typedef struct {} nfc_ndef_data;
// std::vector<uint8_t> page;// TODO: create C vector

// Read Result
typedef struct {
    nfc_info_data info;
    // nfc_pages_data pages;
    // nfc_ndef_data ndef;
    // std::vector<uint8_t> page;
    int status;
} nfc_read_result;


nfc_read_result nfc_read(nfc_read_conf);

#endif