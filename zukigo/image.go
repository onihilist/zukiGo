package zukigo

type ImageRequest struct {
	Prompt  string `json:"prompt"`
	Model   string `json:"model"`
	Size    string `json:"size"`
	Quality string `json:"quality"`
	N       int    `json:"n"`
}

func (z *ZukiInterface) ImageCall(model, prompt, size, quality string, n int) ([]byte, error) {
	reqBody := ImageRequest{
		Prompt:  prompt,
		Model:   model,
		Size:    size,
		Quality: quality,
		N:       n,
	}
	return z.makeRequest("/v1/images/generations", reqBody)
}
