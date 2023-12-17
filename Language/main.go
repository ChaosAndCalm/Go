package main

import (
	"fmt"
)

func main(){
	s := "ēlite"
	fmt.Printf("%8T %[1]v\n",s)

	fmt.Printf("%8T %[1]v\n", []rune(s))
	fmt.Printf("%8T %[1]v\n", []byte(s))

	var str string = "Hello, World!"

	str += "World!"
	fmt.Printf("%d\n",len(str))
}