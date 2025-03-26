# ZukiAPI (Go Wrapper)

ZukiAPI is a lightweight REST API written in Golang that serves as a **gateway (wrapper)** around a third-party OpenAI-compatible API (e.g. [zukijourney.xyzbot.net](https://zukijourney.xyzbot.net)).

It exposes two simple endpoints:
- `/chat` for text generation
- `/image` for image generation

---

## 🚀 Features

✅ REST API built with Gin  
✅ Authenticates requests using an API key  
✅ Clean `ZukiInterface` client structure  
✅ Easily extensible for other endpoints

---

## 🛠️ Tech Stack

- Go 1.20+
- Gin (web framework)
- net/http (standard HTTP client)
- JSON encoding/decoding
- `.env` file support (optional)

---

## ⚙️ Setup

### 1. Clone the repo

```bash
git clone https://github.com/youruser/zukiapi.git
cd zukiapi
```

### 2. Initialize Go modules

```bash
go mod tidy
```

### 3. Set your API key

```bash
export ZUKI_API_KEY=your_api_key_here
```

### 4. Run the API

```bash
go run main.go
```

---

## 🔌 Available Endpoints

`POST /chat`

Request:
```json
{
  "model": "gpt-3.5-turbo",
  "messages": [
    { "role": "user", "content": "Hello!" }
  ]
}
```

Response :

```json
{
  "id": "...",
  "choices": [...]
}
```
---

`POST /image`

Request:
```json
{
  "prompt": "A robot in Paris",
  "model": "flux-schnell",
  "size": "512x512",
  "quality": "standard",
  "n": 1
}
```

Response :

```json
{
  "data": [
    { "url": "https://..." }
  ]
}
```