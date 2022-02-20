package userbiz

import (
	"Fresher_go/common"
	"Fresher_go/modules/user/usermodel"
	"context"
)

type RegisterDb interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}
type Hasher interface {
	Hash(data string) string
}
type registerBussiness struct {
	registerDb RegisterDb
	hasher     Hasher
}

func NewRegisterUserBiz(registerDb RegisterDb, hasher Hasher) *registerBussiness {
	return &registerBussiness{registerDb: registerDb, hasher: hasher}
}

// can ktra xem email nay da ton tai hay chua r moi cho dky
func (bussiness *registerBussiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := bussiness.registerDb.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}
	salt := common.GenSalt(50)
	data.Password = bussiness.hasher.Hash(data.Password + salt)
	data.Status = 1
	data.Salt = salt
	data.Role = "user"
	if err := bussiness.registerDb.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
