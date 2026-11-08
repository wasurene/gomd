package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wasurene/gomd/internal/converter"
	"github.com/wasurene/gomd/internal/storage"
)

func Run() {
	if len(os.Args) <= 1 {
		fmt.Println("Ошибка: напиши команду")
		fmt.Println("Пример команд: go run . new about / go run . build")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		slicebyte := []byte("# " + os.Args[2])
		pathMd := filepath.Join("content", os.Args[2]+".md")

		err := os.MkdirAll("content", 0755)
		if err != nil {
			fmt.Printf("Ошибка при создании папки: %v", err)
		}

		err = os.WriteFile(pathMd, slicebyte, 0644)
		if err != nil {
			fmt.Printf("Ошибка при записи в файл: %v", err)
		}

		fmt.Printf("Файл %v создан", pathMd)

	case "build":
		mdFiles, err := storage.GetMarkdownFiles("content")
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}

		for _, mdFile := range mdFiles {
			pathMd := filepath.Join("content", mdFile)
			pathHTML := filepath.Join("public", strings.TrimSuffix(mdFile, ".md")+".html")

			bytesMd, err := storage.ReadFile(pathMd)
			if err != nil {
				fmt.Printf("Ошибка: %v", err)
			}

			bytesHTML := converter.MarkdownToHTML(bytesMd)

			err = storage.WriteFile(pathHTML, bytesHTML)
			if err != nil {
				fmt.Printf("Ошибка: %v", err)
			}
		}

	default:
		fmt.Println("Неизвестная команда")
	}
}
