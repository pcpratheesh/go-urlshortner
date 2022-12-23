package models

// HTTPError400 BadRequestError
type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPError500 HTTPError
type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"status internal server error"`
}

// HTTPError404 HTTPError NotFound Error
type HTTPError404 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"status Page NotFound"`
}

// HTTPOk StatusOK
type HTTPOk struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"status Ok"`
}

// HTTPCreated StatusAccepted
type HTTPCreated struct {
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"status Ok"`
}
