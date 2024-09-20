package main

import (
    "fmt"
    "strconv"
)

func main() {
    var a string
    fmt.Scan(&a)

    max := 0 
    for _, digit := range a {
        num, _ := strconv.Atoi(string(digit)) // Преобразуем символ в число

        if num > max {
            max = num
        }
    }
    fmt.Print(max)
}