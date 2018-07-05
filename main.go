package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type Resume struct {
	Header Header
	Body   Body
}

type Header struct {
	Name     string
	Github   string
	LinkedIn string
	Title    string
	Email    string
	Phone    string
	Location string
}

type Body struct {
	Works []Work
}

type Work struct {
	Company     string
	Time        string
	Title       string
	Description []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	template := template.Must(template.ParseFiles("resume_template.html"))

	resume, err := os.Create("./resume.html")
	check(err)
	defer resume.Close()

	raw, err := ioutil.ReadFile("./resume.json")
	check(err)
	resumeData := BytesToResume(raw)

	// data := Resume{
	// 	Header: Header{
	// 		Name:     "Vinh Tran",
	// 		Github:   "daivinhtran",
	// 		LinkedIn: "vinhtran40",
	// 	},
	// }
	template.Execute(resume, resumeData)

	// data := TodoPageData{
	// 	PageTitle: "My TODO list",
	// 	Todos: []Todo{
	// 		{Title: "Task 1", Done: false},
	// 		{Title: "Task 2", Done: true},
	// 		{Title: "Task 3", Done: true},
	// 	},
	// }
	// template.Execute(resume, data)

}

func BytesToResume(raw []byte) *Resume {
	resume := &Resume{}
	err := json.Unmarshal(raw, resume)
	check(err)
	return resume
}
