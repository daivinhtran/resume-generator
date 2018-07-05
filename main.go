/*
A simple resume generator from json to html
Author: Vinh Tran
*/
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

	template.Execute(resume, resumeData)
}

func BytesToResume(raw []byte) *Resume {
	resume := &Resume{}
	err := json.Unmarshal(raw, resume)
	check(err)
	return resume
}
