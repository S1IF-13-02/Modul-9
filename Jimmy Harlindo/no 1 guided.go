package main
import "fmt"

func main(){
	var angka int

	fmt.Print("masukan nilai:")
	fmt.Scan(&angka)

	if angka < 0 {
		angka = - angka
	}
fmt.Print("nilai  obsolut:",angka)
}