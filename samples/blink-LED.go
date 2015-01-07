package main

import (
	"fmt"
	"github.com/audreylim/vspark"
	"log"
	"time"
)

func main() {
	err := vspark.PingSpark()
	if err != nil {
		log.Fatal(err)
	}

	err = vspark.PinMode("D7", "OUTPUT")
	if err != nil {
		fmt.Println(err)
	}

	for {
		writeErr := vspark.DigitalWrite("D7", 1)
		if err != nil {
			fmt.Println(writeErr)
		}
		time.Sleep(1000 * time.Millisecond)

		writeErr = vspark.DigitalWrite("D7", 0)
		if err != nil {
			fmt.Println(writeErr)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
