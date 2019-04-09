package main

import (
	"fmt"
	"os"
	sc "strconv"
)

func main() {
	lhs, _ := sc.Atoi(os.Args[2])
	rhs, _ := sc.Atoi(os.Args[3])
	switch os.Args[1] {
	case "add":
		fmt.Println(lhs + rhs)
	case "sub":
		fmt.Println(lhs - rhs)
	}
}
