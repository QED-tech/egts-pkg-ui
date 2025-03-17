package server

type Problem struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type Response struct {
	Data interface{} `json:"data"`
}
