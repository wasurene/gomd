package converter

import "fmt"

func MarkdownToHTML(md []byte) []byte {
	htmlString := fmt.Sprintf("<html><body><pre> %s </pre></body></html>", string(md))

	return []byte(htmlString)
}
