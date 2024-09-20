package main
import "fmt"
import "math"

func main() {
    var a, b float64
    fmt.Scan(&a, &b)
    
    c := (a * a + b * b)
    fmt.Print(math.Sqrt(c))
}