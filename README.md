# go-vinilStore-api
# 📀 Go Vinyl Store API

## 📌 Описание проекта
Go Vinyl Store API — это RESTful API, реализованный на Go с использованием фреймворка **Gin**. API позволяет управлять коллекцией виниловых пластинок, предоставляя CRUD-операции для работы с данными.

## 🚀 Функциональность
- 📋 **Получение списка альбомов** (`GET /albums`)
- 🔍 **Получение альбома по ID** (`GET /albums/:id`)
- ➕ **Добавление нового альбома** (`POST /albums`)

## 🛠️ Установка и запуск

### 1️⃣ **Клонирование репозитория**
```sh
 git clone https://github.com/aventhis/go-vinilStore-api.git
 cd go-vinilStore-api
```

### 2️⃣ **Установка зависимостей**
```sh
 go mod tidy
```

### 3️⃣ **Запуск сервера**
```sh
 go run cmd/main.go
```

Сервер будет запущен по адресу **`http://localhost:8080`**.

## 📡 API Эндпоинты

### 📋 Получение всех альбомов
```http
GET /albums
```
**Пример ответа:**
```json
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    }
]
```

### 🔍 Получение альбома по ID
```http
GET /albums/:id
```
**Пример ответа:**
```json
{
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
}
```

### ➕ Добавление нового альбома
```http
POST /albums
```
**Пример запроса:**
```json
{
    "id": "4",
    "title": "Kind of Blue",
    "artist": "Miles Davis",
    "price": 49.99
}
```
**Пример ответа:**
```json
{
    "message": "Album added successfully",
    "album": {
        "id": "4",
        "title": "Kind of Blue",
        "artist": "Miles Davis",
        "price": 49.99
    }
}
```

## 🛠 Используемые технологии
- [Go](https://go.dev/) — основной язык разработки
- [Gin](https://github.com/gin-gonic/gin) — веб-фреймворк
- [cURL](https://curl.se/) — для тестирования API

## ✨ TODO
- 🔄 Добавить возможность обновления альбома (`PUT /albums/:id`)
- 🗑 Удаление альбома (`DELETE /albums/:id`)
- 📦 Подключить базу данных (PostgreSQL или SQLite)

## 👨‍💻 Автор
[aventhis](https://github.com/aventhis)
