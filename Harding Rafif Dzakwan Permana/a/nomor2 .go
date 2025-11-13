package main

import "fmt"

func main () {
    var n int
    fmt.Print("Masukan nilai n : ")
    fmt.Scan(&n)

    keterangan := "bukan" 

	if n < 0 && n % 2 == 0{
      keterangan = "genap negatif "
	}

	fmt.Print(keterangan)
}
