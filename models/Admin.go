package models

type Admin struct {
	AdminId  string `json:"admin_id" gorm:"primaryKey" gorm:"column:admin_id" validate:"required"`
	Username string `json:"username" gorm:"column:username" validate:"required"`
	Password string `json:"password" gorm:"column:password" validate:"required"`
}
