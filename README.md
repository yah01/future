# Future

Async-Await style for Golang

## Get Started
call a function async:
```golang
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

	fmt.Printf("hello %s", future.Value())
}
```

An async call returns a future, which contains the return value of the method. Invoking `Await()` method to wait for the result.