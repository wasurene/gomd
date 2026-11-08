package storage

import (
	"fmt"
	"os"
	"strings"
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
