package common

import "kasir-api/common/dtos"

func SuccessResponse() dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: true,
			Message: "Successfully"},
	}
}

func SuccessResponseWithData(data interface{}) dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: true,
			Message: "Successfully"},
		Data: data,
	}
}

func GetListSuccessResponse[T any](data []T) dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: true,
			Message: "Successfully",
		},
		Total: len(data),
		Data:  data,
	}
}

func NotFoundResponse() dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: false,
			Message: "Not Found"},
	}
}
