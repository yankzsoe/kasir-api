package dtos

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Response struct {
	Status BaseResponse `json:"status"`
	Data   interface{}  `json:"data,omitempty"`
	Total  int          `json:"total,omitempty"`
}

type PaginationMeta struct {
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	TotalData  int  `json:"total_data"`
	TotalPages int  `json:"total_pages"`
	HasPrev    bool `json:"has_prev"`
	HasNext    bool `json:"has_next"`
}

type PaginationResponse struct {
	Status BaseResponse   `json:"status"`
	Data   interface{}    `json:"data,omitempty"`
	Meta   PaginationMeta `json:"meta"`
}

type ErrorResponse struct {
	ErrorCode int
	Message   Response
}
