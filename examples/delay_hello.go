package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/yah01/future"
)

func DelayJoin(names ...string) string {
	time.Sleep(time.Second)
	result := strings.Join(names, ", ")

	return result
}

func main() {
	future := Submit(func() (any, error) {
		res := DelayJoin("yah01", "zer0ne")
		return res, nil
	})

	fmt.Printf("hello %s", future.Value())
}
