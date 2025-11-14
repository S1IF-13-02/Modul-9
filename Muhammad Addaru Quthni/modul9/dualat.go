package main
import "fmt"

func main() {
	var a int
	var hasil = "positif"
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&a)
	if a <= 0 {
		hasil = "bukan positif"
	}
	fmt.Print(hasil)
}