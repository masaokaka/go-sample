package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

var x = 1

const y = 2

func main() {
	z := 3
	fmt.Println(x, y, z)

	fmt.Println(puppy.Bark(), puppy.BigBark())
	puppy.From13()

}
