package structs

type ValidatorData struct {
	Index   string `json:"index"`
	Balance string `json:"balance"`
	Status string `json:"status"`
}

type Response struct {
	Data []ValidatorData `json:"data"`
}
