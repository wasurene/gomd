package converter

import (
	"bytes"
	"html/template"
)

type PageData struct {
	Title   string
	Content string
	Pages   []string
}

type FrontMatter struct {
	Title string `yaml:"title"`
}

func MarkdownToHTML(data PageData) ([]byte, error) {
	sample, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		return []byte{}, err
	}

	var buff bytes.Buffer

	err = sample.Execute(&buff, data)
	if err != nil {
		return []byte{}, err
	}

	return buff.Bytes(), nil
}
