package storage

import (
	"strings"
	"testing"
)

func TestParseFrontMatter(t *testing.T) {
	input := []byte("---\ntitle: Мой крутой заголовок\n---\n# Привет, мир!")

	front, sliceByte, err := ParseFrontMatter(input)
	if err != nil {
		t.Fatalf("Функция вернула неожиданную ошибку: %v", err)
	}

	if front.Title != "Мой крутой заголовок" {
		t.Errorf("Ожидали 'Мой крутой заголовок', получили '%s'", front.Title)
	}

	stringByte := string(sliceByte)
	stringByte = strings.TrimSpace(stringByte)

	if stringByte != "# Привет, мир!" {
		t.Errorf("Ожидали '# Привет, мир!', получили '%s'", stringByte)
	}

}
