package notestorage

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
)

func (s *sqlStore) FindDataByCondition(
	cxt context.Context,
	conditions map[string]interface{},
	morekeys ...string) (*notemodel.Note, error) {

	var result notemodel.Note

	db := s.db

	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}

	if err := db.Where(conditions).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
