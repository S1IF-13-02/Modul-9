package main
import "fmt"
func main () {

	var x int
	fmt.Scan(&x)

	jumlahmotor := x/2
	if x %2 != 0 {
	jumlahmotor++
	}
	fmt.Print(jumlahmotor)
	
}