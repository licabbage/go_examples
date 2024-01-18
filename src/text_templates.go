package main

import (
	"os"
	"text/template"
)

func main() {

	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value is {{.}}\n")) // {{.}} 表示值

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n") //{{.Name}}表示值的Name属性

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n") //If 语句
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	t3.Execute(os.Stdout, true)
	t3.Execute(os.Stdout, 0)
	t3.Execute(os.Stdout, 1)

	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})

	t4.Execute(os.Stdout, map[int]string{
		1: "a",
		2: "b",
	})
}
