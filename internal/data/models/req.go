package models

type Req struct{
	ID uint64 `json:"id"`
}

type Resp struct{
	Error string `json:"error"`
	Status int `json:"status"`
}