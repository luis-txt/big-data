package main

import "fmt"

func main() {
	data := generateRandomData(10, 0.0, 20.0)
	res := alonMatiasSzegedy(data, 4)
	fmt.Println(res)
}
