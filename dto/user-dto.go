package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	NIK      string `json:"nik" form:"nik" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	NIK      string `json:"nik" form:"nik" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}
