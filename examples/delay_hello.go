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
	future := AsyncCall(func() string {
		return DelayJoin("yah01", "zer0ne")
	})
	result := future.Await()

	fmt.Printf("hello %s", result)
}
