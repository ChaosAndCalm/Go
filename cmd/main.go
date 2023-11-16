package main

import (
	"fmt"
	"hello"
	"os"
)

func main(){
	 fmt.Println(hello.Say(os.Args[1:]))

	 a := 3
	 b := 2.01

	 fmt.Printf("a: %T %v\n",a,a)
	 fmt.Printf("b: %T %v\n",b,b)
}