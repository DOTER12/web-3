package main
import (
    "fmt"
    "strings"
)

func main() {
    var a string
    fmt.Scan(&a)
    
    b := strings.Split(a, "")
    c := strings.Join(b, "*")
    
    fmt.Println(c)
    
}