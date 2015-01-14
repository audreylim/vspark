#Vspark

Vspark is a Go package that lets you write programs for [Spark devices] (https://www.spark.io/).

##Getting Started

Load the [Voodoospark firmware] (https://github.com/voodootikigod/voodoospark) onto your Spark device. 

*For users familiar with [Spark-cli] (https://github.com/spark/spark-cli), voodoospark firmware is bundled and can be flashed through the CLI.

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
    if writeErr != nil {
      fmt.Println(writeErr)
    }
    time.Sleep(1000 * time.Millisecond)

    writeErr = vspark.DigitalWrite("D7", 0)
    if writeErr != nil {
      fmt.Println(writeErr)
    }
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
  err := vspark.PingSpark()
  if err != nil {
    log.Fatal(err)
  }
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
err := PinMode("D7", "OUTPUT")

```

**DigitalWrite(pin, value)**

Set 1 for maximum 3.3V voltage ("HIGH") or 0 to turn it off ("LOW").

Example: 
```go
err := PinMode("D7", "OUTPUT")
writeErr := DigitalWrite("D7", 1)
```

**AnalogWrite(pin, value)**

Available on D0, D1, A0, A1, A4, A5, A6, A7.

Set value from 0 to 255.

Example:
```go
err := PinMode("A1", "PWM")
writeErr := AnalogWrite("A1", 100)
```

**DigitalRead(pin)**

Example:
```go
err := PinMode("D1", "INPUT")
val, readErr := DigitalRead("D1")
```

**AnalogRead(pin)**

Example:
```go
err := PinMode("A1", "ANALOG")
val, readErr := AnalogRead("A1")
```

**ServoWrite(pin, deg)**

Set deg between 0 to 180.

Available on D0, D1, A0, A1, A4, A5, A6, A7.

Example:
```go
err := PinMode("D1", "SERVO")
servoErr := ServoWrite("D1", 180)
```

##LICENSE

MIT
