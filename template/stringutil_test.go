package template

import (
	"bytes"
	"regexp"
	"testing"
	ttemplate "text/template"
)

func TestSuperReplace(t *testing.T) {
	ticket3 := `AK: "xxxxxxxxxxxx"
    SK: "xxxxxxxxxxxxxxxxxxxxxx"
    BaseUrl: "https://xxxx.xxx.xxx"`
	tmpl3 := `AK: "[[[.wps.ak]]]"
	SK: "[[[.wps.sk]]]"
	BaseUrl: "[[[.wps.url]]]"`
	values3 := make(map[string]any)
	values3["wps"] = map[string]string{"ak": "ak", "sk": "sk", "url": "url"}

	result, err := SuperReplace(ticket3, tmpl3, values3)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("================result=================")
	t.Logf("%s,result", result)

}

func TestTemplate(t *testing.T) {
	//	tmpl := `
	//WPS:
	//	[[[.Version]]]
	//    DocsVersion: "[[[.Version]]]"
	//`
	tmpl2 := `sadffs[[.Version]]]asdfasfd`
	template := ttemplate.New("super_replace").Delims("[[[", "]]]")
	p, err := template.Parse(tmpl2)
	if err != nil {
		t.Fatal(err)
	}
	var b bytes.Buffer
	if err := p.Execute(&b, nil); err != nil {
		t.Fatal(err)
	}
	s, e := findStarEndFlag(p.Tree)
	t.Log(s, e, err)
}

func TestName(t *testing.T) {
	str := "abc foo:bar def baz:qux ghi"
	re := regexp.MustCompile("([a-z]+):([a-z]+)")
	result := ReplaceAllStringSubmatchFunc(re, str, func(groups []string) string {
		return groups[1] + "." + groups[2]
	})
	t.Logf("'%s'\n", result)
}
