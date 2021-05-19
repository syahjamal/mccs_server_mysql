package model

type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(225)" json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(225)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Books    *[]Book `json:"books,omitempty"`
}
