package tmpl

// https://blog.gopheracademy.com/advent-2017/using-go-templates/
import (
	"os"
	"strings"
	"testing"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func TestTempl(t *testing.T) {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	template := template.New("todos")
	template, err := template.Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	if err != nil {
		panic(err)
	}
	err = template.Execute(os.Stdout, td)
	if err != nil {
		t.Fatal(err)
	}

	tdNew := Todo{"Go", "Contribute to any Go project"}
	if err = template.Execute(os.Stdout, tdNew); err != nil {
		t.Fatal(err)
	}
}

// Must functions, used to verify that a template is valid during parsing
func TestMust(t *testing.T) {
	tdNew := Todo{"Go", "Contribute to any Go project"}
	r := template.Must(template.New("todos").Parse("You have task named \"{{ .Name}}\" with description: \"{{ .Description}}\""))
	var sb strings.Builder
	if err := r.Execute(&sb, tdNew); err != nil {
		t.Fatal(err)
	}
	t.Log(sb.String())
}

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func TestHtmlTodo(t *testing.T) {
	// Files are provided as a slice of strings.
	paths := []string{
		"~/Develop/go_practice/generate/tmpl/todo.tmpl",
	}
	todos := ToDo{
		User: "test1",
		List: []entry{
			{Name: "entry1", Done: false},
			{Name: "entry2", Done: true},
		},
	}
	wd, _ := os.Getwd()
	t.Log(wd)
	// template 的名字在与 ParseFiles 一起使用时不是随意取的，务必要与模板文件名字相同。
	//ParseFiles 支持解析多个文件，如果是传入多个文件该咋办？godoc 说了，template 名字与第一个文件名相同即可
	r := template.Must(template.New("todo.tmpl").ParseFiles(paths...))
	var sb strings.Builder
	err := r.Execute(&sb, todos)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sb.String())
}
