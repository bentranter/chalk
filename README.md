# Chalk

Colour your terminal strings. Inspired by Node's [Chalk](https://github.com/chalk/chalk)

### Features

1. The eight default terminal colours

If you want more features, like bold colours or windows support, open an issue or send a PR!

### Usage

Get the package:

```bash
$ go get -u gitub.com/bentranter/chalk
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
        (chalk.Black("Black")),
        (chalk.Red("Red")),
        (chalk.Yellow("Yellow")),
        (chalk.Green("Green")),
        (chalk.Blue("Blue")),
        (chalk.Magenta("Magenta")),
        (chalk.Cyan("Cyan")),
        (chalk.White("White")),
    )
}
```

### License

MIT &copy; 2015 Ben Tranter
