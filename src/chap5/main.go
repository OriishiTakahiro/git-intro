package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// ログの出力先を指定
	logfile, err := os.OpenFile("./logs/"+time.Now().Format("2006-01-02_15-04-05.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	log.SetOutput(logfile)
	// "Hello, <入力文字>"をログに残して表示
	msg := "Hello, " + os.Args[1] + "!"
	log.Println(msg)
	fmt.Println(msg)
}
