package main

import (
	"os"
	"text/template"
)

func main() {
	tEmpty := template.New("template test")
	//empty pipeline following if
	tEmpty.Execute(os.Stdout, nil)
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n")) 

	tWithValue := template.New("template test")
	//non empty pipeline following if condition
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	//non empty pipeline following if condition
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}