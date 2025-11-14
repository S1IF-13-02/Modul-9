package main

import "fmt"

func main() {
    var orang int
	var motor int
	
    fmt.Print("Masukkan jumlah orang: ")
    fmt.Scan(&orang)

    motor = orang / 2

    if orang%2 == 1 {
        motor = motor + 1
    }

    fmt.Println(motor)
}