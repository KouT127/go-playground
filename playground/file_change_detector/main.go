package main

import (
	"fmt"
	"time"
)

type directory struct {
	files []filename
}

type filename struct {
	name string
}

func main() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		}
	}
}
