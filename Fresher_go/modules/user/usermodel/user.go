package usermodel

import (
	"Fresher_go/common"
	"errors"
)

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"-" gorm:"password"`
	LastName        string        `json:"last_name" gorm:"last_name"`
	FirstName       string        `json:"first_name" gorm:"first_name"`
	Role            string        `json:"-" gorm:"role"`
	Salt            string        `json:"-" gorm:"salt"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

const EntityName = "User"

//create
type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"password" gorm:"password"`
	LastName        string        `json:"last_name" gorm:"last_name"`
	FirstName       string        `json:"first_name" gorm:"first_name"`
	Role            string        `json:"-" gorm:"role"`
	Salt            string        `json:"-" gorm:"salt"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"email;"`
	Password string `json:"password" form:"password" gorm:"password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}
func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
