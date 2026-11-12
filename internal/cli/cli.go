package cli

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/wasurene/gomd/internal/config"
	"github.com/wasurene/gomd/internal/converter"
	"github.com/wasurene/gomd/internal/storage"
)

func Run() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("Ошибка при чтении конфига: %v", err)
		os.Exit(1)
	}

	if len(os.Args) <= 1 {
		fmt.Println("Ошибка: напиши команду")
		fmt.Println("Пример команд: go run . new about / go run . build")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		file := filepath.Join(cfg.TemplateDir, "default.md")
		def, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Ошибка при чтении шаблона: %v", err)
			os.Exit(1)
		}

		result := strings.ReplaceAll(string(def), "{{.Title}}", os.Args[2])
		pathMd := filepath.Join(cfg.ContentDir, os.Args[2]+".md")

		err = os.MkdirAll(cfg.ContentDir, 0755)
		if err != nil {
			fmt.Printf("Ошибка при создании папки: %v", err)
		}

		err = os.WriteFile(pathMd, []byte(result), 0644)
		if err != nil {
			fmt.Printf("Ошибка при записи в файл: %v", err)
		}

		fmt.Printf("Файл %v создан", pathMd)

	case "build":
		var sliceMd []string

		mdFiles, err := storage.GetMarkdownFiles(cfg.ContentDir)
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}

		for _, mdFile := range mdFiles {
			sliceMd = append(sliceMd, strings.TrimSuffix(mdFile, ".md"))
		}

		for _, mdFile := range mdFiles {

			pathMd := filepath.Join(cfg.ContentDir, mdFile)
			pathHTML := filepath.Join(cfg.OutputDir, strings.TrimSuffix(mdFile, ".md")+".html")

			bytesMd, err := storage.ReadFile(pathMd)
			if err != nil {
				fmt.Printf("Ошибка: %v", err)
			}

			fm, clearMd, err := storage.ParseFrontMatter(bytesMd)
			if err != nil {
				fmt.Printf("Ошибка: %v", err)
			}

			data := converter.PageData{
				Title:   fm.Title,
				Content: template.HTML(string(clearMd)),
				Pages:   sliceMd,
			}

			templatePath := filepath.Join(cfg.TemplateDir, "layout.html")
			bytesHTML, err := converter.MarkdownToHTML(data, templatePath)
			if err != nil {
				fmt.Printf("Ошибка: %v", err)
			}

			err = storage.WriteFile(pathHTML, bytesHTML)
			if err != nil {
				fmt.Printf("Ошибка: %v", err)
			}
		}
	case "serve":
		fileServer := http.FileServer(http.Dir(cfg.OutputDir))
		err := http.ListenAndServe(":8080", fileServer)
		if err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Println("Неизвестная команда")
	}
}
