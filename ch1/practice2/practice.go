package main

import (
	"fmt"
	"os"
)

func main() {
	for index, args := range os.Args[1:] {
		fmt.Printf("i: %d,args: %s \n", index, args)
	}
}
