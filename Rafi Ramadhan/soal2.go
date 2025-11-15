package main
import "fmt"

func main() {
    var angka int
	var hasil string
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&angka)

    hasil = "bukan"         
    if angka%2 == 0 && angka < 0 {
        hasil = "genap negatif"
    }

    fmt.Println(hasil)
}
