package main

import (
	"bcd2/bcd"
	"fmt"
	"log"
	"os"
)

func main() {
	x, err := bcd.Aton(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	y, err := bcd.Aton(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	w := bcd.Add(x, y)
	fmt.Printf("%v + %v = %v\n", x, y, w)
}
