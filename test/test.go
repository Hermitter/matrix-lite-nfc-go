package main

import (
	"fmt"
	"time"

	nfc "github.com/matrix-io/matrix-lite-nfc-go"
)

func main() {
	for {
		status := nfc.Status(0)
		fmt.Println(status)

		time.Sleep(50 * time.Millisecond)
	}
}
