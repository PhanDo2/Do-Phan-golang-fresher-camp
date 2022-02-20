package restaurantdb

import (
	"Fresher_go/common"
	"Fresher_go/modules/restaurents/restaurantmodel"
	"context"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

//create
func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.Restaurants) error {
	//câu lệnh truy vấn (query statement)
	if err := s.db.Model(&restaurantmodel.Restaurants{}).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

//search
func (s *sqlStore) Find(cxt context.Context,
	conditions map[string]interface{},
	morekeys ...string) (*restaurantmodel.Restaurants, error) {
	var result restaurantmodel.Restaurants
	for i := range morekeys {
		s.db.Preload(morekeys[i])
	}

	if err := s.db.Where(conditions).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

//update
func (s *sqlStore) Update(cxt context.Context,
	id int,
	data *restaurantmodel.Restaurants) error {

	//if err := s.db.Model(&restaurantmodel.Restaurants{}).Where("id=?", id).UpdateColumn("status", data.Status).Error; err != nil {
	//	return err
	//}
	// để như này k update được dữ liệu của status, city_id
	if err := s.db.Where("id=?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

//list data
func (s *sqlStore) List(cxt context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	morekeys ...string) ([]restaurantmodel.Restaurants, error) {

	//var result []restaurantmodel.Restaurants
	//
	//for i := range morekeys {
	//	s.db = s.db.Preload(morekeys[i])
	//}
	//db := s.db.Table("restaurants").Where(conditions)
	//
	//if err := db.Count(&paging.Total).Error; err != nil {
	//	return nil, common.ErrDB(err)
	//}
	//if paging.FackCursor != "" {
	//	if uid, err := common.FromBase58(paging.FackCursor); err != nil {
	//		db = db.Where("id < ?", uid.GetLocalID())
	//	} else {
	//		db = db.Offset((paging.Page - 1) * paging.Page)
	//	}
	//
	//}
	//if err := s.db.Limit(paging.Limit).Find(&result).Error; err != nil {
	//
	//	return nil, common.ErrDB(err)
	//}
	//return result, nil
	var result []restaurantmodel.Restaurants

	db := s.db

	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}

	db = db.Table(restaurantmodel.Restaurants{}.TableName()).Where(conditions).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FackCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

//delete
func (s *sqlStore) Delete(
	ctx context.Context,
	id int,
) error {

	if err := s.db.Table(restaurantmodel.Restaurants{}.TableName()).
		Where("id=?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrCannotDeleteEntity("data", err)
	}
	return nil
}
