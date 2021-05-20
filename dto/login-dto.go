package dto

type LoginDTO struct {
	NIK      string `json:"nik" form:"nik" binding:"required"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
