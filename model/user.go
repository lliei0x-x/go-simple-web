package model

import (
	"fmt"
	"time"

	"go-simple-web/common"
	"go-simple-web/config"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	db     = config.POSTGRES
	secret = config.GetServerSecret()
)

const (
	avatarURL = "https://www.gravatar.com/avatar/%s?d=retro" // https://en.gravatar.com/site/implement/images
	lastSeen  = "last_seen"
	aboutMe   = "about_me"
	pwdHash   = "password_hash"
)

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

// SetPasswordHash func
func (u *User) SetPasswordHash(pwd string) {
	u.PasswordHash = common.GeneratePasswordHash(pwd)
}

// CheckPassword func
func (u *User) CheckPassword(pwd string) bool {
	return common.GeneratePasswordHash(pwd) == u.PasswordHash
}

// GetUserByUsername func
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail func
func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// SetAvatar func
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf(avatarURL, common.Md5(email))
}

// GenerateToken func
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 可以添加过期时间
	})
	return token.SignedString([]byte(secret))
}

// CheckToken func
func CheckToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}
	return "", err
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
	contents := map[string]interface{}{lastSeen: time.Now()}
	return UpdateUserByUsername(username, contents)
}

// UpdateAboutMe func
func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{aboutMe: text}
	return UpdateUserByUsername(username, contents)
}

// UpdatePassword func
func UpdatePassword(username, password string) error {
	contents := map[string]interface{}{pwdHash: common.Md5(password)}
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
