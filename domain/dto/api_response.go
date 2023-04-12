package dto

// ApiResponse is a generic response for all API endpoints
type ApiResponse struct {
	Code    int         `json:"code" example:"200"`
	Status  string      `json:"status" example:"OK"`
	Message string      `json:"message" example:"Your request has been processed successfully"`
	Data    interface{} `json:"data" swaggertype:"object" nullable:"true"`
}
