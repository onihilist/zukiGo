package zukigo

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

func (z *ZukiInterface) ChatCall(model string, messages []ChatMessage) ([]byte, error) {
	reqBody := ChatRequest{
		Model:    model,
		Messages: messages,
	}
	return z.makeRequest("/v1/chat/completions", reqBody)
}
