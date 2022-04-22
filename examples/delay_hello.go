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
	future := Submit(func() (string, error) {
		res := DelayJoin("yah01", "zer0ne")
		return res, nil
	})

	fmt.Printf("hello %s\n", future.Value())

	future1 := Submit(func() (int, error) {
		time.Sleep(200 * time.Millisecond)
		return 5, nil
	})

	future2 := Submit(func() (string, error) {
		time.Sleep(200 * time.Millisecond)
		return "hello", nil
	})

	AwaitAll(future1, future2)

	fmt.Println(future1.Value(), future2.Value())
}
