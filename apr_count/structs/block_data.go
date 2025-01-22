package structs

type Message struct {
	Slot           string `json:"slot"`
}

type Data struct {
	Msg *Message `json:"message"`
}

type Slot struct {
	BlockData *Data `json:"data"`
}
