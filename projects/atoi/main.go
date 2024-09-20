package main
import(
    "fmt"
)

func main() {
    
    var a string
    fmt.Scan(&a)
    
    for i := 0; i < len(a); i++ {
    tmp := a[i] - 48
    fmt.Print(tmp * tmp)
    }
    
}