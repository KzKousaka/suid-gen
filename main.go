package main

import (
	"flag"
	"fmt"

	"github.com/lithammer/shortuuid/v4"
)

func main() {
	var length = 0
	var num = 0

	flag.IntVar(&length, "l", 22, "length of short-uuid")
	flag.IntVar(&num, "n", 100, "number of short-uuids generated")
	flag.Parse()

	if length > 22 {
		length = 22
	}

	for i := 0; i < num; i++ {
		u := shortuuid.New()
		fmt.Println(u[0:length])
	}
}
