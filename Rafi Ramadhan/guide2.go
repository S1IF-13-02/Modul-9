package main
import "fmt"
func main() {
    var angka int
	 var ispositif bool
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&angka)
    if angka > 0 {
        ispositif = true
    }

    fmt.Println("hasil: ", ispositif)
}
