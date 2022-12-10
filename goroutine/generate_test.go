package goroutine

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	repeat := func(
		done <-chan interface{}, values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{}, valueStream <-chan interface{}, num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case x := <-valueStream:
					takeStream <- x
				}
			}
		}()
		return takeStream
	}
	done := make(chan interface{})
	defer close(done)
	for num := range take(done, repeat(done, 1), 10) {
		t.Logf("%v ", num)
	}
}
