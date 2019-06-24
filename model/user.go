package model

import (
	"fmt"
	"time"

	"go-simple-web/common"
	"go-simple-web/config"
)

var db = config.POSTGRES

// User struct
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(64)"`
	PasswordHash string `gorm:"type:varchar(64)"`
	LastSeen     time.Time
	AboutMe      string `gorm:"type:varchar(128)"`
	Avatar       string `gorm:"type:varchar(200)"`
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
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// SetAvatar func
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=retro", common.Md5(email)) // https://en.gravatar.com/site/implement/images
}

// UpdateUserByUsername func
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	u, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(u).Update(contents).Error
}

// UpdateLastSeen func
func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

// UpdateAboutMe func
func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, contents)
}

// AddUser func
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPasswordHash(password)
	user.SetAvatar(email)
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return user.FollowSelf()
}
