package user

type Buyer struct {
	BuyID              int    `json:"buyID"`
	TelNum             string `json:"telNum"`
	RealName           string `json:"realName"`
	UserName           string `json:"userName"`
	Gender             string `json:"gender"`
	MarkAdd            string `json:"markAdd"`
	EmailAddress       string `json:"emailAddress"`
	CollShopIDs        string `json:"collShopIDs"`
	CollProductIDs     string `json:"collProductIDs"`
	Password           string `json:"password"`
	IdentificationType string `json:"identificationType"`
	IdentificationNum  string `json:"identificationNum"`
	Region             string `json:"region"`
	Birthday           string `json:"birthday"`
	AliFlag            string `json:"aliFlag"`
	RecomID            int    `json:"recomID"`
	RecomTime          string `json:"recomTime"`
	BankAccount        string `json:"bankAccount"`
}
