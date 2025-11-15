package main 
import "fmt"
func main() {
	 var n int
    fmt.Print("Masukkan sebuah bilangan bulat: ")
    fmt.Scan(&n)
if n < 0 {
    n = -n
}
    fmt.Println("Nilai absolut adalah:", n)
}
