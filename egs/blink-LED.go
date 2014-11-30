package main

import (
	"github.com/audreylim/vspark"
	"time"
)

func main() {
	vspark.PingSpark()
	vspark.PinMode("D7", "OUTPUT")

	for {
		vspark.DigitalWrite("D7", 1)
		time.Sleep(1000 * time.Millisecond)

		vspark.DigitalWrite("D7", 0)
		time.Sleep(1000 * time.Millisecond)
	}
}
