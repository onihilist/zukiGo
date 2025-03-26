package main

import (
	"fmt"
	"github.com/onihilist/zukigo"
)

func main() {
	// Define your API key here
	zuki := zukigo.New("your_api_key")

	messages := []zukigo.ChatMessage{
		{Role: "user", Content: "Hello world"},
	}

	// Call for generate a chat
	resp, err := zuki.ChatCall("gpt-3.5-turbo", messages)
	if err != nil {
		fmt.Println("Chat error:", err)
	} else {
		fmt.Println("Chat response:", string(resp))
	}

	// Call for generate an image
	imgResp, err := zuki.ImageCall("flux-schnell", "A robot in Paris", "512x512", "standard", 1)
	if err != nil {
		fmt.Println("Image error:", err)
	} else {
		fmt.Println("Image response:", string(imgResp))
	}

}
