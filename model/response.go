package model

// Response adalah struct response umum.
type Response struct {
	Message string      `json:"message" example:"detail pesan"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"detail error"`
}

// UnauthorizedResponse adalah contoh response untuk 401 Unauthorized.
type UnauthorizedResponse struct {
	Message string `json:"message" example:"token tidak valid atau tidak ditemukan"`
}

// ForbiddenResponse adalah contoh response untuk 403 Forbidden.
type ForbiddenResponse struct {
	Message string `json:"message" example:"user tidak memiliki akses untuk fitur ini"`
}

// SuccessResponse adalah contoh response untuk 200 OK.
type SuccessResponse struct {
	Message string      `json:"message" example:"berhasil"`
	Data    interface{} `json:"data,omitempty"`
}

// CreatedResponse adalah contoh response untuk 201 Created.
type CreatedResponse struct {
	Message string      `json:"message" example:"data berhasil ditambahkan"`
	Data    interface{} `json:"data,omitempty"`
}
