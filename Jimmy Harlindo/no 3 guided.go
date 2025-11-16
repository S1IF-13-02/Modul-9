package main
import "fmt"

func main(){
	var bilangan int
	var hasil bool

	fmt.Print("masukan bilangan")
	fmt.Scan(&bilangan)

	hasil = false
	if bilangan < 0 && bilangan%2 == 0{
		hasil = true
	}
fmt.Println(hasil)
}