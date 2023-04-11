package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID             uint      `gorm:"primarykey"`
	UserName       string    `gorm:"column:user_name;type:varchar(255)" json:"user_name"`
	PasswordDigest string    `gorm:"column:password_digest;type:longtext" json:"password_digest"`
	Nickname       string    `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
	Status         string    `gorm:"column:status;type:varchar(255)" json:"status"`
	Avatar         string    `gorm:"column:avatar;type:varchar(1000)" json:"avatar"`
	UpdatedAt      time.Time `json:"update_at" db:"update_at" description:"更新时间"`
	CreatedAt      time.Time `json:"create_at" db:"create_at" description:"创建时间"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
