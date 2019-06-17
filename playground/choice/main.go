package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//　引数
	c := len(os.Args) - 1
	if c < 1 {
		//　出力先指定でFormatを指定できる。
		fmt.Fprintf(os.Stderr, "[usage] %s choce1 choice2...", os.Args[0])
		// プログラムを終了できる
		os.Exit(1)
	}
	fmt.Printf(os.Args[rand.Intn(c)+1])
}
