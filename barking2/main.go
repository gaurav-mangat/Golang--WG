package main

import (
	"fmt"

	"github.com/gaurav-mangat/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(s3)
	fmt.Println(s4)
}