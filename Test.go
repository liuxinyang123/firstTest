package main

import "fmt"

type a struct {
	num int
}

func main() {
	var p = &a{3}
	p = nil
	fmt.Println(p)
}
