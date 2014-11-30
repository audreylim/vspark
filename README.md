#Vspark

Vspark is a Go package that lets you write Go programs that interact with Spark devices.

##Getting Started

Load the [Voodoospark firmware] (https://github.com/voodootikigod/voodoospark) (v 2.4.0)  onto your Spark device. 

Add the following to your dot file:

```go
export SPARK_TOKEN="your spark token"
export SPARK_DEVICE_ID="your device id"
```

##Install

```go
$ go get github.com/audreylim/vspark
```

##Use

**blink-LED.go**

```go
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
```

##API

**StaticIP and PingSpark()**

Set StaticIP if you want to assign a static IP address, or just use `PingSpark()` to get the IP address of the Spark Core.

Eg.
```go 
func main() {
  vspark.StaticIP = "your static IP address and port number"
  vspark.PingSpark()
}
```  

**PinMode(pin, mode)**

Modes:
- INPUT:           0 (DigitalRead, AnalogRead)
- OUTPUT:          1 (DigitalWrite)
- PWM:             1 (AnalogWrite)
- INPUT_PULLUP:    2
- INPUT_PULLDOWN:  3
- SERVO:           4 (ServoWrite)

Eg. 
```go
PinMode("D7", "OUTPUT")
```

**DigitalWrite(pin, value)**

Set 1 for maximum 3.3V voltage ("HIGH") or 0 to turn it off ("LOW").

Eg. 
```go
PinMode("D7", "OUTPUT")
DigitalWrite("D7", 1)
```

**AnalogWrite(pin, value)**

Available on D0, D1, A0, A1, A4, A5, A6, A7.

Set value from 0 to 255.

Eg.
```go
PinMode("A1", "PWM")
AnalogWrite("A1", 100)
```

**DigitalRead**

Eg.
```go
PinMode("D1", "INPUT")
DigitalRead("D1")
```

**AnalogRead**

Eg.
```go
PinMode("D1", "INPUT")
AnalogRead("D1")
```

**ServoWrite**

Set value between 0 to 180.

Available on D0, D1, A0, A1, A4, A5, A6, A7.

Eg.
```go
PinMode("D1", "SERVO")
ServoWrite("D1", 180)
```

##LICENSE

MIT
