# Chalk

Colour your terminal strings. Inspired by Node's [Chalk](https://github.com/chalk/chalk)

### Usage

Get the package:

```bash
$ go get -u github.com/bentranter/chalk
```

Use it:

```go
package main

import (
  "fmt"
  "github.com/bentranter/chalk"
)

func main() {
  // Look at the colours!
  fmt.Println(
    (Black("Black")),
    (Red("Red")),
    (Yellow("Yellow")),
    (Green("Green")),
    (Blue("Blue")),
    (Magenta("Magenta")),
    (Cyan("Cyan")),
    (White("White")),
  )
}
```

### License

MIT &copy; 2015 Ben Tranter
