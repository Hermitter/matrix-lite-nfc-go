package main

import (
	"fmt"
	"time"

	matrix "github.com/matrix-io/matrix-lite-go"
	nfc "github.com/matrix-io/matrix-lite-nfc-go"
)

func main() {
	m := matrix.Init()
	fmt.Println("Starting NFC Read")

	for {
		code, _ := nfc.Read(nfc.ReadConf{true, false, false, 0})

		if code == 256 {
			m.Led.Set(matrix.RGBW{0, 0, 1, 0})
		} else if code == 1024 {
			m.Led.Set("black")
		}

		fmt.Println(code, nfc.Status(code))
		// fmt.Printf("%+v\n", tag)

		time.Sleep(1 * time.Millisecond)
	}
}
