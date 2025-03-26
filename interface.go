package zukigo

type ZukiInterface struct {
	Auth    string
	BaseURL string
}

func New(auth string) *ZukiInterface {
	return &ZukiInterface{
		Auth:    auth,
		BaseURL: "https://api.zukijourney.com",
	}
}
