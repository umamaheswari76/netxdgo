package main

import "fmt"

func main() {
    var i interface{}
    i = 42 // Assign an int value to the interface

    // Type assertion to extract the int value
    if val, ok := i.(int); ok {
        fmt.Println("Successfully asserted:", val)
    } else {
        fmt.Println("Type assertion failed")
    }
}