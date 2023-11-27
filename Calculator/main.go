package main

import (
	"fmt"
)

func main(){
	a := 2
	fmt.Printf("a: %8T %[1]v\n",a)

	b := 2.1
	fmt.Printf("b: %8T %[1]v\n",b)
	
}