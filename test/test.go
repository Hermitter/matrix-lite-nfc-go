package main

import (
	"fmt"

	nfc "github.com/matrix-io/matrix-lite-nfc-go"
)

func main() {
	status := nfc.Status(0)

	fmt.Println(status)
}
