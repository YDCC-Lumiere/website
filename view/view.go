package view

import (
	"html/template"
	"bytes"
	"net/http"
	"log"

  "github.com/gin-gonic/gin"
)

type View struct {
	TempHTML []byte
	Tmpl *template.Template
}

type AlertData struct {
	Type string
	Title string
	Content string
}

var view View

func Initialize(glob string) {
	view.TempHTML = nil
	tmpl, err := view.Tmpl.ParseGlob(glob)
	if err != nil {
		log.Printf("Init template error: %v", err)
	}
	view.Tmpl = tmpl
	log.Printf("Init templates, %v", view.Tmpl.Templates())
}

func Alert(_type string, title string, content string) {
	Add("alert.tmpl", AlertData{Type: _type, Title: title, Content: content})
}

func Add(name string, data any) {
	buf := new(bytes.Buffer)

	err := view.Tmpl.ExecuteTemplate(buf, name, data)

	if err != nil {
		log.Printf("[Add] err = %v", err)
	}

	log.Printf("[Add] buf = %v", string(buf.Bytes()[:]))

	view.TempHTML = buf.Bytes()
}

func Text(c *gin.Context, text string) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(text))
}

func Execute(c *gin.Context, name string, data any) {
	buf := new(bytes.Buffer)

	err := view.Tmpl.ExecuteTemplate(buf, name, data)

	if err != nil {
		log.Printf("[Execute] err = %v", err)
	}

	log.Printf("[Execute] buf = %v", string(buf.Bytes()[:]))

	c.Data(http.StatusOK, "text/html; charset=utf-8", append(buf.Bytes(), view.TempHTML...))
	view.TempHTML = nil
}
