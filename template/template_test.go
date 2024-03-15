package template

import (
	"os"
	"testing"
	ttemplate "text/template"
	"text/template/parse"
)

type Person struct {
	Version string
}

func TestBBB(t *testing.T) {
	t.Skip()
	tmpl := `
WPS:
    DocsVersion: "[[[.Version]]]"
	CC: "{{var.etcd_endpoints_1}}"
`
	//p := Person{Version: "xsdf"}
	p := map[string]string{"Version": "111111"}
	t3 := ttemplate.New("test")
	t3.Delims("[[[", "]]]")
	parseInstance, err := t3.Parse(tmpl)
	if err != nil {
		t.Error(err)
	}
	t2 := parseInstance.Tree
	t.Logf("root pos: %d", t2.Root.Position())
	for i, n := range t2.Root.Nodes {
		t.Log()
		t.Logf("%d,%d,%t,%s", i, n.Position(), n.Type() == parse.NodeText, n)
	}
	p2 := parseInstance.Root.Position()

	t.Log(p2)
	err = parseInstance.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
