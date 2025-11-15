package main
import "fmt"

func main() {
    var x, y int
	   var faktorXY, faktorYX bool
    fmt.Print("Masukkan dua bilangan bulat (x y): ")
    fmt.Scan(&x, &y)

    if true { 
        faktorXY = y%x == 0
        faktorYX = x%y == 0
    }
    fmt.Println(faktorXY)
    fmt.Println(faktorYX)
}
