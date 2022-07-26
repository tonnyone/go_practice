package types

import (
	"fmt"
)

//go:generate ./newList2.sh Person

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.Age)
}
