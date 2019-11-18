package main

import (
	"fmt"
	"time"

	nfc "github.com/matrix-io/matrix-lite-nfc-go"
)

func main() {
	fmt.Println("Starting NFC Read")

	for {
		code, tag := nfc.Read(nfc.ReadConf{true, false, false, 0})

		if code == 256 {
			fmt.Printf("%+v\n", tag)
		} else if code == 1024 {
			fmt.Println("...Nothing Scanned...")
		}

		time.Sleep(1 * time.Millisecond)
	}
}
