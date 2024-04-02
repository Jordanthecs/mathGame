package main

import(
	"fmt"
	"math/rand"
)

func main(){
	var x int
	var y int
	x = rand.Intn(20)+1
	y = rand.Intn(20)+1
	fmt.Print(x ," + ", y);
	fmt.Print(" = ",x+y)
}