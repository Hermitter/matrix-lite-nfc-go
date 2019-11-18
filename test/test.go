package main

import (
	"fmt"
	matrix "github.com/matrix-io/matrix-lite-go"
	nfc "github.com/matrix-io/matrix-lite-nfc-go"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting NFC Read")

	wg.Add(1)
	read(&wg)

	// for {
	// 	fmt.Println("Hello!")
	// 	time.Sleep(5 * time.Millisecond)
	// }

	wg.Wait()

}

func read(wg *sync.WaitGroup) {
	m := matrix.Init()
	for {
		code, tag := nfc.Read(nfc.ReadConf{true, false, false, 30})

		if code == 256 {
			fmt.Printf("%+v\n", tag)
			// m.Led.Set("blue")
		} else if code == 1024 {
			fmt.Println("...Nothing Scanned...")
			m.Led.Set("black")
		}

		time.Sleep(1 * time.Millisecond)
	}

	wg.Done()
}
