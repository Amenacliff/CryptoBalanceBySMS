package dtos

type GetBalanceReqDTO struct {
	PhoneNumber string `json:"phoneNumber"`
	Token       string `json:"token"`
}
