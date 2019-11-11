package main

import (
	"fmt"
	"time"

	nfc "github.com/matrix-io/matrix-lite-nfc-go"
)

func main() {
	for {
		code, tag := nfc.Read(nfc.ReadConf{true, false, false, 0})
		fmt.Println(code, nfc.Status(code))
		fmt.Printf("%+v\n", tag)

		time.Sleep(50 * time.Millisecond)
	}
}
