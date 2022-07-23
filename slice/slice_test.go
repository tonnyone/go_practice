package slice

import (
	"testing"
)

func TestSlice(t *testing.T) {
	a := make([]int, 32)
	// a := []int{0, 0}
	// t.Logf("a: %v, %d, %d", a, len(a), cap(a))
	b := a[1:16]
	t.Logf("a: %v, %d, %d", a, len(a), cap(a))
	a = append(a, 1) // append后,会变成新的地址，copy了
	t.Logf("a: %v, %d, %d", a, len(a), cap(a))
	t.Logf("b: %v, %d, %d", b, len(b), cap(b))
	t.Logf("a: %p", a)
	t.Logf("b: %p", b)
}

func TestAppend(t *testing.T) {
	a := make([]int, 32)
	t.Logf("a: %p, %v, %d, %d", a, a, len(a), cap(a))
	t.Logf("a[0] address: %p", &a[0])
	a = append(a, 1)
	t.Logf("a[0] address: %p", &a[0])
	t.Logf("a: %p, %v, %d, %d", a, a, len(a), cap(a))
}

// TestSlice1 非常经典的一道题
func TestSlice1(t *testing.T) {
	sl := make([]int, 0, 10)
	var appFn = func(s []int) {
		s = append(s, 10, 20, 30)
		t.Log(s)
	}

	t.Log(sl)               // 输出
	appFn(sl)               // 输出
	t.Log(len(sl), cap(sl)) // 输出
	t.Log(sl)               // 输出
	t.Log(sl[:])            // 输出
	t.Log(sl[:10])          // 输出
	sl[4] = 50
}

// TestSlice1 非常经典的一道题
func TestSliceMem(t *testing.T) {
	sl := make([]int, 0, 10)
	t.Logf("%p", sl)
	var appFn = func(s []int) {
		t.Logf("%p", s)
		s = append(s, 10, 20, 30)
		t.Logf("%p", s)
	}
	appFn(sl)
	t.Logf("%p", sl)
	sl = append(sl, 40)
	t.Logf("%p", sl)
	sl = append(sl, 50)
	t.Logf("%p", sl)
}
