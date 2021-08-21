package main
import (
	"html/template"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDos struct {
	User string
	List []entry
}

func main()  {
	paths := []string{
		"todo.tmpl",
	}


	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
	todos := ToDos{
		User: "Emmanuel",
		List: []entry{},
	}
	err := t.Execute(os.Stdout, todos)
	if err != nil{
		panic(err)
	}
}