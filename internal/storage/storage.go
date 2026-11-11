package storage

import (
	"fmt"
	"os"
	"strings"

	"github.com/wasurene/gomd/internal/converter"
	"gopkg.in/yaml.v3"
)

func WriteFile(path string, data []byte) error {
	err := os.MkdirAll("public", 0755)
	if err != nil {
		fmt.Printf("Ошибка при создании папки: %v", err)
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Printf("Ошибка при записи в файл: %v", err)
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v", err)
		return []byte{}, err
	}

	return file, nil
}

func GetMarkdownFiles(dir string) ([]string, error) {
	filenames := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Ошибка при чтении директории: %v", err)
		return []string{}, err
	}

	for i := range files {
		if !files[i].IsDir() && strings.HasSuffix(files[i].Name(), ".md") {
			filenames = append(filenames, files[i].Name())
		}
	}

	return filenames, nil
}

func ParseFrontMatter(md []byte) (converter.FrontMatter, []byte, error) {
	text := string(md)

	if !strings.HasPrefix(text, "---") {
		fmt.Println("Где открывающие ---?")
		return converter.FrontMatter{}, []byte(md), nil
	}

	text = text[3:]

	endIndex := strings.Index(text, "---")
	if endIndex == -1 {
		fmt.Println("Где закрывающие ---?")
		return converter.FrontMatter{}, []byte(md), nil
	}

	yamlPart := text[:endIndex]

	markdowmPart := text[endIndex+3:]

	var fm converter.FrontMatter

	err := yaml.Unmarshal([]byte(yamlPart), &fm)
	if err != nil {
		return converter.FrontMatter{}, []byte{}, err
	}

	return fm, []byte(markdowmPart), nil
}
