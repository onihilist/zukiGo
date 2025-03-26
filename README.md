# ZukiGo
[![Go Reference](https://pkg.go.dev/badge/github.com/onihilist/zukigo.svg)](https://pkg.go.dev/github.com/onihilist/zukigo)

**ZukiGo** is a lightweight Go client library that acts as a wrapper for a third-party OpenAI-compatible API.

## ✨ Features

- Simple, clean interface for sending chat and image requests  
- Authenticated via Bearer token  
- Fully decoupled: no web server included  
- Easy to integrate into any Go project  


## 🛠 Tech Stack

- **Go 1.20+**
- Standard `net/http` client
- JSON (encoding/decoding)
- Modular Go packages (`interface.go`, `chat.go`, `image.go`, `request.go`)



## 📦 Installation

```bash
go get github.com/onihilist/zukigo@v0.1.3
```


## 🧱 Library Structure

- **interface.go** – Struct and constructor for ZukiInterface

- **chat.go** – ChatCall(model, messages) method

- **image.go** – ImageCall(model, prompt, size, quality, n) method

- **request.go** – Internal method to send HTTP POST requests

## 🔐 Authentication

The library uses a Bearer token for authentication.<br>
`zuki := zukiapi.New("your_api_key")
`

## 🧑‍💻 Usage Example

```go
package main

import (
    "fmt"
    "github.com/onihilist/zukigo"
)

func main() {
    zuki := zukiapi.New("your_api_key")

    // Chat example
    messages := []zukiapi.ChatMessage{
        {Role: "user", Content: "Hello world"},
    }

    // Chat generation exemple
    chatResp, err := zuki.ChatCall("gpt-3.5-turbo", messages)
    if err != nil {
        fmt.Println("Chat error:", err)
        return
    }
    fmt.Println("Chat response:", string(chatResp))

    // Image generation example
    imgResp, err := zuki.ImageCall("flux-schnell", "A robot in Paris", "512x512", "standard", 1)
    if err != nil {
        fmt.Println("Image error:", err)
        return
    }
    fmt.Println("Image response:", string(imgResp))
}
```