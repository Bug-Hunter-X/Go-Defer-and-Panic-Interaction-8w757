```go
package main

import "fmt"

func main() {
    var x int
    fmt.Println(x) // Output: 0
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    //Handle the panic more gracefully, or better, prevent it.
    //Example
    if true {

        panic("This is a panic!")
    }else{
         fmt.Println("This will be printed")
    }

}
```