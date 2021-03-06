# Chalk

[![GoDoc](https://godoc.org/github.com/bentranter/chalk?status.svg)](https://godoc.org/github.com/bentranter/chalk)

[![Build Status](https://semaphoreci.com/api/v1/bentranter/chalk/branches/master/badge.svg)](https://semaphoreci.com/bentranter/chalk)

Colour your terminal strings. Inspired by Node's [Chalk](https://github.com/chalk/chalk)

### Features

1. The eight default terminal colours

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
        (chalk.Black("Black")),
        (chalk.Red("Red")),
        (chalk.Yellow("Yellow")),
        (chalk.Green("Green")),
        (chalk.Blue("Blue")),
        (chalk.Magenta("Magenta")),
        (chalk.Cyan("Cyan")),
        (chalk.White("White")),
    )

    // You can change the background colour too!
    fmt.Println(
        (chalk.BgBlue("My background is blue."))
    )

    // Or underline words
    fmt.Println(
        (chalk.Underline("I'm underlined!"))
    )
}
```

### License

Licensed under Apache 2.0. Copyright 2016-2017 Ben Tranter. See license.md.
