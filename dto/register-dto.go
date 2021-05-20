package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	NIK      string `json:"nik" form:"nik" binding:"required"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
