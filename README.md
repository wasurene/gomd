# GOMD (Go Markdown)

Простой генератор статических сайтов на Go

## Возможности 
- Генерация HTML из Markdown (с помощью goldmark).
- Поддержка YAML Front Matter (метаданные в начале файла).
- Использование HTML-шаблонов для дизайна.
- CLI-команды: new, build, serve.
- Локальный HTTP-сервер для предпросмотра.
- Гибкая настройка путей через `config.yaml`.  

## Установка
```bash
git clone https://github.com/wasurene/gomd
cd gomd
go install ./cmd/gomd
```

## Использование

`gomd new about` — создает файл `content/about.md` с базовым шаблоном.

`gomd build` — собирает все страницы в папку `public`.

`gomd serve` — запускает сервер на `http://localhost:8080`.

## Конфигурация
Все пути настраиваются через один файл `config.yaml`

По умолчанию там
```yaml
content_dir: "content"
output_dir: "public"
template_dir: "templates"
```

## Структура проекта

```text
gomd/
├── cmd/gomd/        # Точка входа
├── internal/        # Внутренняя логика (cli, config, converter, storage)
├── templates/       # HTML-шаблоны и шаблон новой страницы
├── content/         # Исходные Markdown-файлы
└── config.yaml      # Конфигурация
```