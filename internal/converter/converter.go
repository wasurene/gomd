package converter

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
)

type PageData struct {
	Title   string
	Content template.HTML
	Pages   []string
}

type FrontMatter struct {
	Title string `yaml:"title"`
}

func MarkdownToHTML(data PageData, templatePath string) ([]byte, error) {
	md := goldmark.New()
	var buf bytes.Buffer

	err := md.Convert([]byte(data.Content), &buf)
	if err != nil {
		return []byte{}, err
	}

	data.Content = template.HTML(buf.String())

	sample, err := template.ParseFiles(templatePath)
	if err != nil {
		return []byte{}, err
	}

	var templateBuff bytes.Buffer

	err = sample.Execute(&templateBuff, data)
	if err != nil {
		return []byte{}, err
	}

	return templateBuff.Bytes(), nil
}
