package error_my

import (
	"fmt"
	"runtime/debug"
)

// 参考此包的错误处理：
func test1() error {
	fmt.Println("test1")
	stack := debug.Stack()
	return fmt.Errorf("test1 error: %s", stack)
}

func test2() error {
	err := test1()
	if err != nil {
		return fmt.Errorf("test1 error: %v", err)
	}
	return nil
}

func Test3() error {
	err := test2()
	if err != nil {
		return fmt.Errorf("test2 error: %v", err)
	}
	return nil
}
