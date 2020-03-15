package lib

type Response struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data"`
}

