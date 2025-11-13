package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    isXFactor := (y % x) == 0

    isYFactor := (x % y) == 0

    fmt.Println(isXFactor)
    fmt.Println(isYFactor)
}
