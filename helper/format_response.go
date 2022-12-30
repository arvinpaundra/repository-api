package helper

import "net/http"

type Paginate struct {
	Page       int64 `json:"page,omitempty"`
	Limit      int64 `json:"limit,omitempty"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int64 `json:"total_pages"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Page    Paginate    `json:"pagination,omitempty"`
}

// 201 - Created
func SuccessCreatedResponse() BaseResponse {
	return BaseResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: "CREATED",
	}
}

// 200 - OK
func SuccessResponse(data interface{}, paginate Paginate) BaseResponse {
	return BaseResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "OK",
		Data:    data,
		Page:    paginate,
	}
}

// 400 - Bad Request
func BadRequestResponse(errors interface{}) BaseResponse {
	return BaseResponse{
		Code:    http.StatusBadRequest,
		Status:  "error",
		Message: "BAD_REQUEST",
		Errors:  errors,
	}
}

// 401 - Unauthorized
func UnauthorizedResponse() BaseResponse {
	return BaseResponse{
		Code:    http.StatusUnauthorized,
		Status:  "error",
		Message: "UNAUTHORIZED",
	}
}

// 403 - Forbidden
func ForbiddenResponse() BaseResponse {
	return BaseResponse{
		Code:    http.StatusForbidden,
		Status:  "error",
		Message: "FORBIDDEN",
	}
}

// 404 - Not Found
func NotFoundResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusNotFound,
		Status:  "error",
		Message: message,
	}
}

// 409 - Conflict
func ConflictResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusConflict,
		Status:  "error",
		Message: message,
	}
}

// 415 - Unsupported Media Type
func UnsupportedMediaTypeResponse(errors interface{}) BaseResponse {
	return BaseResponse{
		Code:    http.StatusUnsupportedMediaType,
		Status:  "error",
		Message: "UNSUPPORTED_MEDIA_TYPE",
		Errors:  errors,
	}
}

// 500 - Internal Server Error
func InternalServerErrorResponse(message string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusInternalServerError,
		Status:  "error",
		Message: message,
	}
}
