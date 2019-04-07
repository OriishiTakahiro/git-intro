package main

import (
	"fmt"
	"os"
	sc "strconv"
)

func main() {
	lhs, lerr := sc.Atoi(os.Args[2])
	rhs, rerr := sc.Atoi(os.Args[3])
	if lerr != nil || rerr != nil {
		fmt.Println("2, 3番目の引数の何れかが整数に変換できません")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		fmt.Println(lhs + rhs)
	case "sub":
		fmt.Println(lhs - rhs)
	case "div":
		fmt.Println(lhs / rhs)
	}
}
