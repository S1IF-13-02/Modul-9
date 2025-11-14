package main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Masukkan dua angka: ")
    fmt.Scanln(&x, &y)
    if true {
        fmt.Println(y % x == 0 , x % y == 0)
    }
}