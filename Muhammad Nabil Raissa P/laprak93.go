package main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Masukkan dua bilangan (x y): ")
    fmt.Scan(&x, &y)

    if x != 0 && y%x == 0 {
        fmt.Println("true")
    }
    if x == 0 || y%x != 0 {
        fmt.Println("false")
    }

    if y != 0 && x%y == 0 {
        fmt.Println("true")
    }
    if y == 0 || x%y != 0 {
        fmt.Println("false")
    }
}
