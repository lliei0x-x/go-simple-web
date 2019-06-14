package model

import (
	"go-simple-web/common"
	"go-simple-web/config"
)

// User struct
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

// SetPasswordHash ...
func (u *User) SetPasswordHash(pwd string) {
	u.PasswordHash = common.GeneratePasswordHash(pwd)
}

// CheckPassword ...
func (u *User) CheckPassword(pwd string) bool {
	return common.GeneratePasswordHash(pwd) == u.PasswordHash
}

// GetUserByUsername ...
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := config.POSTGRES.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
