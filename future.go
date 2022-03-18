package future

import (
	"runtime"
)

type Empty = struct{}

type ConcurrencyLimiter struct {
	limit chan Empty
}

func NewConcurrencyLimiter(limit int) ConcurrencyLimiter {
	return ConcurrencyLimiter{
		limit: make(chan struct{}, limit),
	}
}

func (limiter *ConcurrencyLimiter) Acquire() {
	limiter.limit <- Empty{}
}

func (limiter *ConcurrencyLimiter) Release() {
	<-limiter.limit
}

func (limiter *ConcurrencyLimiter) Len() int {
	return len(limiter.limit)
}

func (limiter *ConcurrencyLimiter) Cap() int {
	return cap(limiter.limit)
}

var (
	globalLimiter = NewConcurrencyLimiter(runtime.GOMAXPROCS(0))
)

func ReplaceGlobalLimiter(limit int) {
	globalLimiter = NewConcurrencyLimiter(limit)
}

type Future[T any] struct {
	value   chan T
	limiter ConcurrencyLimiter
}

func (future Future[T]) Await() T {
	return <-future.value
}

func withGlobalLimiter[T any]() Future[T] {
	return Future[T]{
		value:   make(chan T),
		limiter: globalLimiter,
	}
}

// func AsyncCall[T any](method func() T) Future[T] {
// 	future := withGlobalLimiter[T]()
// 	future.limiter.Acquire()

// 	go func() {
// 		defer future.limiter.Release()
// 		value := method()
// 		future.value <- value
// 	}()

// 	return future
// }

func AsyncCall[T any](method func() T, limiters ...*ConcurrencyLimiter) Future[T] {
	var future Future[T]
	if len(limiters) == 0 {
		future = withGlobalLimiter[T]()
	} else {
		future = Future[T]{
			value:   make(chan T),
			limiter: *limiters[0],
		}
	}
	future.limiter.Acquire()

	go func() {
		defer future.limiter.Release()
		value := method()
		future.value <- value
	}()

	return future
}

// func AsyncCall2[T any, Arg1 any, Arg2 any](method func(arg1 Arg1, arg2 Arg2) T, arg1 Arg1, arg2 Arg2) Future[T] {
// 	future := withGlobalLimiter[T]()
// 	future.limiter.Acquire()

// 	go func() {
// 		defer future.limiter.Release()
// 		value := method(arg1, arg2)
// 		future.value <- value
// 	}()

// 	return future
// }

// func AsyncCall3[T any, Arg1 any, Arg2 any, Arg3 any](method func(arg1 Arg1, arg2 Arg2, arg3 Arg3) T,
// 	arg1 Arg1, arg2 Arg2, arg3 Arg3) Future[T] {
// 	future := withGlobalLimiter[T]()
// 	future.limiter.Acquire()

// 	go func() {
// 		defer future.limiter.Release()
// 		value := method(arg1, arg2, arg3)
// 		future.value <- value
// 	}()

// 	return future
// }
