package main

import (
	"fmt"
	"os"
	"text/template"
)

func useTemplatesAsString() {
	fmt.Println("Using template as string")
	temp1 := template.New("testTemplate")
	temp1, err := temp1.Parse("{{ range . }}Hi {{ .Name}} please meet my friend {{ .FriendName }}\n{{end}}") //parsing template as string and range is used to iterate over the data passed to template
	if err != nil {
		fmt.Println("Error while parsing template", err)
	}

	err = temp1.Execute(os.Stdout, []map[string]string{{"Name": "Pradyum", "FriendName": "Karthik"}, {"Name": "Pradyum", "FriendName": "Ashir"}}) //providing list for the template to iterate over it.
	if err != nil {
		fmt.Println("Error while executing template", err)
	}

}

func useTemplateFromFile() {

	fmt.Println("Using template from file")
	temp1, err := template.ParseFiles("templateSample.txt")
	if err != nil {
		fmt.Println("Error while loading template file", err)
	}

	err = temp1.Execute(os.Stdout, map[string]string{"Name": "Pradyum", "FriendName": "Karthik"}) //providing list for the template to iterate over it.
	if err != nil {
		fmt.Println("Error while executing template", err)
	}

}

func main() {

	useTemplatesAsString()
	useTemplateFromFile()

}
