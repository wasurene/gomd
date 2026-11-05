package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run() {
	if len(os.Args) <= 1 {
		fmt.Println("Ошибка: напиши команду")
		fmt.Println("Пример команд: go run . new about / go run . build")
		os.Exit(1)
	}
	fmt.Println("Программа запущена, вы написали: ", os.Args[1])

	switch os.Args[1] {
	case "new":
		fmt.Println("Создаю новую страницу")
		slicebyte := []byte("# about")
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
		fmt.Println("Собираю сайт")
	default:
		fmt.Println("Неизвестная команда")
	}
}
