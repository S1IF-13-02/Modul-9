package main
import "fmt"
func main() {
	var b int
	var n string = "bukan"

	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&b)
	if b<0 && b%2 == 0 {
		n = "genap negatif"
	}
	fmt.Println(n)
}