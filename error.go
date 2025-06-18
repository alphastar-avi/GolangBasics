package main

import (
    "errors"
    "fmt"
)

// Add two numbers, return error if any is negative
func add(a, b int) (int, error) {
    if a < 0 || b < 0 {
        return 0, errors.New("negative numbers are not allowed")
    }
    return a + b, nil
}

func main() {
    result, err := add(5, -3)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Sum:", result)
}
