package vspark

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

//
// Get Spark IP and Ping Spark.
//

var conn *net.TCPConn
var StaticIP string

func PingSpark() {
	// Get Spark Core IP Address.
	var sparkIP string

	if StaticIP != "" {
		sparkIP = StaticIP
	} else {
		reqUrl := fmt.Sprintf(
			"https://api.spark.io/v1/devices/%s/endpoint?access_token=%s",
			os.Getenv("SPARK_DEVICE_ID"),
			os.Getenv("SPARK_TOKEN"),
		)

		client := &http.Client{}
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
			return
		}

		resp, requestErr := client.Do(req)
		if requestErr != nil {
			log.Fatal("Do: ", requestErr)
			return
		}
		defer resp.Body.Close()

		body, dataReadErr := ioutil.ReadAll(resp.Body)
		if dataReadErr != nil {
			log.Fatal("ReadAll: ", dataReadErr)
			return
		}

		type endpoint struct {
			Result string
		}

		var r endpoint
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Fatal(err)
		}

		sparkIP = r.Result
	}

	// Ping Spark with TCP.
	tcpAddr, err := net.ResolveTCPAddr("tcp4", sparkIP)
	if err != nil {
		panic(err)
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
}

//
// PinMode and Pin Commands.
//

// Convert Pin number to its byte represetation.
var n = map[string]byte{
	"D1": 1, "A1": 1,
	"D2": 2, "A2": 2,
	"D3": 3, "A3": 3,
	"D4": 4, "A4": 4,
	"D5": 5, "A5": 5,
	"D6": 6, "A6": 6,
	"D7": 7, "A7": 7,
}

// Available Pins for PWM and Servo writes.
// PWM should be written with AnalogWrite, Servo with ServoWrite.
var availablePins = []string{
	"A0",
	"A1",
	"A4",
	"A5",
	"A6",
	"A7",
	"D0",
	"D1",
}

func PinMode(pin, mode string) {
	m := map[string]byte{
		"INPUT":  0,
		"OUTPUT": 1, "PWM": 1, // Use "OUTPUT" for DigitalWrite, "PWM" for AnalogWrite. TODO: Implement coerced mode?
		"INPUT_PULLUP":   2,
		"INPUT_PULLDOWN": 3,
		"SERVO":          4,
	}

	if string(pin[0]) == "A" {
		n[pin] += 10
	}

	if mode == "SERVO" || mode == "PWM" {
		for _, p := range availablePins {
			if p == pin {
				_, err := conn.Write([]byte{0x00, n[pin], m[mode]})
				if err != nil {
					fmt.Println(err)
				}
				return
			}
		}
		fmt.Printf("%s is not available on %s.\n", mode, pin)
		return
	}

	_, err := conn.Write([]byte{0x00, n[pin], m[mode]})
	if err != nil {
		fmt.Println(err)
	}
}

func DigitalWrite(pin string, value byte) {
	_, err := conn.Write([]byte{0x01, n[pin], value})
	if err != nil {
		fmt.Println(err)
	}
}

func AnalogWrite(pin string, value byte) {
	for _, p := range availablePins {
		if p == pin {
			_, err := conn.Write([]byte{0x02, n[pin], value})
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}

	fmt.Printf("PWM is not available on %s.\n", pin)
}

// Read bytes sent by Voodoo.
func readBytes() uint16 {
	a := [][]byte{}

	// Digital and Analog read actions on Voodoo return 4 bytes: action, pin, lsb, msb.
	for i := 0; i < 4; i++ {
		b := make([]byte, 1)

		_, err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		a = append(a, b)
	}

	// Join 7-bit lsb, msb.
	lsb16 := uint16(a[2][0])
	msb16 := uint16(a[3][0])
	return lsb16 | msb16<<7
}

func DigitalRead(pin string) uint16 {
	// FIXME: https://github.com/voodootikigod/voodoospark/issues/36
	// 3rd byte (1 or 0 only) needed for initial Digital and Analog Read to work, but should be removed afterwards. Not documented.
	_, err := conn.Write([]byte{0x03, n[pin], 1})
	if err != nil {
		fmt.Println(err)
	}
	return readBytes()
}

func AnalogRead(pin string) uint16 {
	_, err := conn.Write([]byte{0x04, n[pin], 1})
	if err != nil {
		fmt.Println(err)
	}

	// NOTE: Didn't seem necessary to implement support for compatibility with firmata's 10-bit ADC values, as Spark-IO did.
	return readBytes()
}

func AlwaysSendBit(pin, value string) {
	v := map[string]byte{
		"DIGITAL": 1,
		"ANALOG":  2,
	}

	_, err := conn.Write([]byte{0x05, n[pin], v[value]})
	if err != nil {
		fmt.Println(err)
	}
}

func ServoWrite(pin string, deg byte) {
	for _, p := range availablePins {
		if p == pin {
			_, err := conn.Write([]byte{0x41, n[pin], deg})
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}

	fmt.Printf("ServoWrite is not available on %s.\n", n[pin])
}

// TODO: Implement SERIAL, SPI, and WIRE actions when ready on Voodoo.
