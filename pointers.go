package main

import "fmt"

func main() {
    number := 42
    fmt.Println(number)  // Output: 42

    pointer := &number
    fmt.Println(pointer) // Output: memory address of 'number'

    *pointer = 100
    fmt.Println(&number)  // Output: 100
}