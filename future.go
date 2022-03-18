package future

import "runtime"

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

var (
	globalLimiter = NewConcurrencyLimiter(runtime.GOMAXPROCS(0))
)

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

func AsyncCall[T any, Arg any](method func(arg Arg) T, arg Arg) Future[T] {
	future := withGlobalLimiter[T]()
	future.limiter.Acquire()

	go func() {
		defer future.limiter.Release()
		value := method(arg)
		future.value <- value
	}()

	return future
}
