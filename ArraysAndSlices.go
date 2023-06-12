package main

import "fmt"

func main() {
    // Arrays
    var numbers [3]int
    numbers[0] = 1
    numbers[1] = 2
    numbers[2] = 3
    fmt.Println(numbers)

    // Slices
    var fruits []string
    fruits = append(fruits, "apple")
    fruits = append(fruits, "banana", "cherry")
    fmt.Println(fruits)
}
