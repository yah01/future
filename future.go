package future

import (
	"github.com/panjf2000/ants/v2"
)

type Future[T any] struct {
	ch    chan struct{}
	value T
	err   error
}

func NewFuture[T any]() *Future[T] {
	return &Future[T]{
		ch: make(chan struct{}),
	}
}

func (future *Future[T]) Await() (T, error) {
	<-future.ch

	return future.value, future.err
}

func (future *Future[T]) Value() T {
	<-future.ch

	return future.value
}

func (future *Future[T]) Err() error {
	<-future.ch

	return future.err
}

func (future *Future[T]) OK() bool {
	<-future.ch

	return future.err == nil
}

func (future *Future[T]) Inner() <-chan struct{} {
	return future.ch
}

type Pool struct {
	inner *ants.Pool
}

func NewRuntime(cap int, opts ...ants.Option) *Pool {
	pool, err := ants.NewPool(cap, opts...)
	if err != nil {
		panic(err)
	}
	return &Pool{
		inner: pool,
	}
}

func (r *Pool) Submit(method func()) error {
	return r.inner.Submit(method)
}
