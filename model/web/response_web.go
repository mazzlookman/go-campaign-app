package web

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type APIResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
