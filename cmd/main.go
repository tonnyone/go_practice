package main

import (
	"fmt"
	"github.com/tonnyone/go_practice/error_my"
)

func main() {
	err := error_my.Test3()
	if err != nil {
		fmt.Println(err)
	}
}
