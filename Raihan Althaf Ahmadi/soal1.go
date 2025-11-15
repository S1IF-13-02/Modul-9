package main
import "fmt"
func main () {
	var x, y int
	fmt.Print("Masukan Jumlah anggota touring : ")
	fmt.Scan(&x)
	y = x/2
	if x%2 != 0{
		y = y + 1
	}
	fmt.Println("Motor yang di perlukan untuk touring adalah : ", y)
}