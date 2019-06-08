package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var ary []int
	var max = 200.0
	//　初期値
	for i := 2; i <= int(max); i++ {
		ary = append(ary, i)
	}
	//　処理部
	bf_t := time.Now()
	ary = eratosthenes(ary, max)
	af_t := time.Now()
	// 表示
	for _, val := range ary {
		print(val)
		print("　")
	}
	fmt.Println(af_t.Sub(bf_t))
}

func eratosthenes(ary []int, max float64) []int {
	for _, val := range ary {
		if int(math.Sqrt(max)) < val {
			break
		}
		ary = remove(&ary, val)
	}
	return ary
}

func remove(ary *[]int, div int) []int {
	var filtered []int
	for _, val := range *ary {
		if val != div && val%div == 0 {
			continue
		}
		filtered = append(filtered, val)
	}
	return filtered
}
