package generate

import (
	"fmt"
	"github.com/tonnyone/go_practice/generate/types"
	"testing"
)

// https://juejin.cn/post/6844903923166216200
// https://docs.google.com/document/d/1V03LUfjSADDooDMhe-_K59EgpTEm3V8uvQRuNMAEnjg/edit#heading=h.j6dsjy94dn2q

// https://ehrt74.medium.com/go-generate-89b20a27f7f9

// 执行方法
// go generate -x

//go:generate ./command.sh

func TestGen(t *testing.T) {
	fmt.Println("if you type 'go generate' in this directory command.sh will be run")
}

//go:generate ./types/newList.sh Person

func TestGen1(t *testing.T) {
	var pl types.PersonList
	pl = append(pl, &types.Person{Name: "Jane", Age: 32})
	pl = append(pl, &types.Person{Name: "Ed", Age: 27})
	pl2 := pl.Filter(func(p *types.Person) bool {
		return p.Age > 30
	})
	for _, p := range pl2 {
		fmt.Println(p)
	}
}
