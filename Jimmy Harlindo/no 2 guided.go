package main
import "fmt"

func main(){
	var bilangan int

	fmt.Print("masukan bilangan")
	fmt.Scan(&bilangan)

	hasil := "bukan positif"
	if bilangan > 0 {
		hasil = "positif"
	}
fmt.Println(hasil)

}