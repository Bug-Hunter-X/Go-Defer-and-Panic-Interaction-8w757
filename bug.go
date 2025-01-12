```go
package main

import "fmt"

func main() {
    var x int
    fmt.Println(x)
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("This is a panic!")
    fmt.Println("This will not be printed")
}
```