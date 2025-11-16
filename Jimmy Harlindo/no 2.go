package main
import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	hasil := "bukan" 

	if bilangan%2 == 0 && bilangan < 0 {
		hasil = "genap negatif"
	}

	fmt.Println(hasil)
}
