package respons

type SuccessResult struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ErrorResult struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
