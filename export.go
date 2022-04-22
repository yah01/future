package future

import "runtime"

var (
	globalPool = NewRuntime(runtime.NumCPU())
)

func Submit[T any](method func() (T, error)) *Future[T] {
	return SubmitWithPool(method, globalPool)
}

func SubmitWithPool[T any](method func() (T, error), pool *Pool) *Future[T] {
	future := NewFuture[T]()
	err := pool.Submit(func() {
		future.value, future.err = method()
		close(future.ch)
	})
	if err != nil {
		future.err = err
		close(future.ch)
	}

	return future
}
