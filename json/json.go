package json

import (
	"encoding/json"
	"fmt"
)

func JsonStr[T any](t T) ([]byte, error) {
	str, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("json marshal error %v", t)
	}
	return str, nil
}

func JsonObj[T any](b []byte) (t T, err error) {
	return t, json.Unmarshal(b, &t)
}

type Person struct {
	Name string
	Age  int
}

type School struct {
	Name    string
	Address string
}

type Body interface {
	*Person | *School
}

type Response[body Body] struct {
	Code int
	Msg  string
	Body body
}
