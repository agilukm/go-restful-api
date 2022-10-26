package response

type Response struct {
	Meta   interface{} `json:"meta"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
	Status string      `json:"status"`
}
