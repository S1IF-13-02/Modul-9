package main 
import "fmt"

 func main() { 
    var x int
    fmt.Print(" masukan jumlah penumpang motor: ")
    fmt.Scan(&x)
    
    if x >= 2 {
        x = (x + 1) / 2
    }
    fmt.Print(x)



    }
