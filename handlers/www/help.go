package web

import (
	"bytes"
	"html/template"
	"net/http"
)

func Help(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("templates/main.html")

	type Params struct {
		Uri     string
		Content string
	}
	p := Params{}

	p.Uri = r.RequestURI

	var b bytes.Buffer
	content, _ := template.ParseFiles("templates" + p.Uri + ".html")
	content.Execute(&b, p)
	p.Content = b.String()

	err := page.Execute(w, p)
	if err != nil {
		return
	}
}
