#Vspark

Vspark is a Go package that lets you write Go programs that interact with [Spark devices] (https://www.spark.io/).

##Getting Started

Load the [Voodoospark firmware] (https://github.com/voodootikigod/voodoospark) onto your Spark device. 

Set these environment variables on the host machine.

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

If you know the IP address of the core and want to bypass the Spark Cloud, you can directly set the IP address. If the IP address is not set, `PingSpark()` will retrieve it from the Spark Cloud.

Example:
```go 
func main() {
  vspark.StaticIP = "your static IP address and port number"
  vspark.PingSpark()
}
```  

**PinMode(pin, mode)**

Modes:
- INPUT => DigitalRead (INPUT_PULLDOWN)
- OUTPUT => DigitalWrite (OUTPUT)
- ANALOG => AnalogRead (INPUT)
- PWM => AnalogWrite (OUTPUT)
- SERVO => ServoWrite (OUTPUT)

Example: 
```go
PinMode("D7", "OUTPUT")
```

**DigitalWrite(pin, value)**

Set 1 for maximum 3.3V voltage ("HIGH") or 0 to turn it off ("LOW").

Example: 
```go
PinMode("D7", "OUTPUT")
DigitalWrite("D7", 1)
```

**AnalogWrite(pin, value)**

Available on D0, D1, A0, A1, A4, A5, A6, A7.

Set value from 0 to 255.

Example:
```go
PinMode("A1", "PWM")
AnalogWrite("A1", 100)
```

**DigitalRead(pin)**

Example:
```go
PinMode("D1", "INPUT")
DigitalRead("D1")
```

**AnalogRead(pin)**

Example:
```go
PinMode("A1", "ANALOG")
AnalogRead("A1")
```

**ServoWrite(pin, deg)**

Set deg between 0 to 180.

Available on D0, D1, A0, A1, A4, A5, A6, A7.

Example:
```go
PinMode("D1", "SERVO")
ServoWrite("D1", 180)
```

##LICENSE

MIT
