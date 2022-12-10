package goroutine

import (
	"testing"
	"time"
)

func TestCloseChannel(t *testing.T) {
	c := make(chan int)
	a := []int{1, 2, 3, 4}
	go func() {
		for _, aa := range a {
			c <- aa
		}
		close(c)
	}()
	for i := 1; i <= 5; i++ {
		r, ok := <-c
		t.Log(i, ":", r, ok)
	}
}

func TestChannelTimeout(t *testing.T) {
	done := make(chan any)
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
	t.Log("loop after")

loop:
	for {
		select {
		case <-done:
			continue loop
		default:
			t.Logf("default: %d", workCounter)
		}
		workCounter++
		time.Sleep(time.Second)
	}
	t.Log(workCounter)
}
