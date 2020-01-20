package main

import (
	"sync"
)

var nums = [5]int64{
	11,
	11,
	11,
	11,
	11,
}

type Total struct {
	number int64
	mutex  sync.Mutex
}

func (t *Total) increment(wg *sync.WaitGroup, num int64) {
	defer wg.Done()
	t.mutex.Lock()
	print(num, "\n")
	t.number += num
	t.mutex.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	total := Total{}
	for _, num := range nums {
		wg.Add(1)
		go total.increment(&wg, num)
	}
	wg.Wait()
	print(total.number, "\n")
}
