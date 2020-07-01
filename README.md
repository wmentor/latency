Simple calc latency library

# Summary

* Very simple
* Written on tiny Go
* Require Go version >= 1.8
* No external dependencies

# Install

```
go get github.com/wmentor/latency
```

# Usage

```
package main

import (
  "fmt"
  "time"

  "github.com/wmentor/latency"  
)

func main() {

  l := latency.New()

  time.Sleep(time.Millisecond * 10)

  fmt.Printf("latency=%.3f\n", l.Seconds()) // like 0.01...

  time.Sleep(time.Millisecond * 160)

  fmt.Printf("latency=%.3f\n", l.Seconds()) // like 0.17...

  l.Duration() // time.Duration
}
```
