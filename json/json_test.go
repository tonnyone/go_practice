package json

import (
	"encoding/json"
	"testing"
)

func TestJsonObj(t *testing.T) {
	s := `{"name":"aaa", "age": 18}`
	obj, err := JsonObj[Person]([]byte(s))
	if err != nil {
		t.Errorf("%v", obj)
	}
	t.Log(obj.Name)
}

func TestJsonGeneric(t *testing.T) {
	resP := &Response[*Person]{
		Code: 200,
		Msg:  "Success",
		Body: &Person{
			Name: "Tim",
			Age:  18,
		},
	}
	str, err := JsonStr[Response[*Person]](*resP)
	if err != nil {
		t.Errorf("json str err")
	}
	t.Log(string(str))
	pJson, err := json.Marshal(resP)
	if err != nil {
		t.Errorf("json marshal error %v", resP)
	}
	xm := &Response[*Person]{}
	if json.Unmarshal(pJson, xm); err != nil {
		t.Errorf("un marshal err")
	}
	var p Person
	p = *(xm.Body)

	t.Log(string(pJson), p)
}
