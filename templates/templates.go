package main

import (
	"os"
	"text/template"
)

type ToDo struct {
	Name	string
	Description	string
}




func main()  {
	td := ToDo{
		Name:        "Test templates",
		Description: "Let's test a template to see the magic",
	}

	t, err := template.New("todos").Parse("You have a task named \"{{.Name}}\" with description: \"{{.Description}}\"")
	if err != nil{
		panic(err)
	}

	err = t.Execute(os.Stdout, td)
	if err != nil{
		panic(err)
	}

	tdNew := ToDo{
		Name:        "Go",
		Description: "Contribute to any Go project",
	}
	errTwo := t.Execute(os.Stdout,tdNew)
	if err != nil{
		panic(errTwo)
	}
}