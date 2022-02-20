package userdb

import (
	"Fresher_go/common"
	"Fresher_go/modules/user/usermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	db := s.db.Table("users")
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	//trả về pointer để có thể trả về nil mà nếu có thì sẽ tái sử dụng lại vùng nhớ (var users usermodel.User) để k bị copy
	var users usermodel.User
	if err := db.Where(conditions).First(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDB(err)
		}
		return nil, common.ErrDB(err)
	}
	return &users, nil
}
