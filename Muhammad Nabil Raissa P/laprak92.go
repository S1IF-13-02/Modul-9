package main 
import "fmt"

 func main() { 
    var x int
    var teks string
    fmt.Print(" masukan bilangan: ")
    fmt.Scan(&x)

    teks = "genap negative"
    
    if x >= 0 {
        teks = "positiv"

  
    }
    fmt.Print(teks)
}
