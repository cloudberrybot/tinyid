# tinyid
Tiny id generator written in Go.

## Installation

```bash
go get github.com/cloudberrybot/tinyid
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/cloudberrybot/tinyid"
)

func main() {
    message := tinyid.Generate(10)
    fmt.Println(message)
}
```
