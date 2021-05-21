package model

// type User struct {
// 	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
// 	Name     string  `gorm:"type:varchar(225)" json:"name"`
// 	Email    string  `gorm:"uniqueIndex;type:varchar(225)" json:"email"`
// 	Password string  `gorm:"->;<-;not null" json:"-"`
// 	Token    string  `gorm:"-" json:"token,omitempty"`
// 	Books    *[]Book `json:"books,omitempty"`
// }

// type MccsUser struct {
// 	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
// 	NIK      string `gorm:"type:varchar(225)" json:"nik"`
// 	Name     string `gorm:"type:varchar(225)" json:"name"`
// 	Password string `gorm:"->;<-;not null" json:"-"`
// 	Token    string `gorm:"-" json:"token,omitempty"`
// }

type MccsUser struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	NIK      string `json:"nik"`
	Name     string `json:"name"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
