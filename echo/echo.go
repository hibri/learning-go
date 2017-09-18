package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	for index := 1; index < len(os.Args); index++ {
		fmt.Println("Index %v %v", index, os.Args[index])
	}
}
