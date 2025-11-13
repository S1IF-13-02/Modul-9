package main

import "fmt"

func main() {
    var orang, motor int

    fmt.Print("Masukkan jumlah orang : ")
    fmt.Scan(&orang)

    motor = (orang + 1) / 2

    fmt.Println("Jumlah motor yang diperlukan :", motor)
}
