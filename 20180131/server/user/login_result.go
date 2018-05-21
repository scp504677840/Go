package user

type LoginResult struct {
	ResMsg string `json:"resMsg"`
	Root   Buyer  `json:"root"`
}
