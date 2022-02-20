package notestorage

import (
	"Fresher_go/common"
	"Fresher_go/modules/notes/notemodel"
	"context"
)

func (s *sqlStore) ListDataByCondition(cxt context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	morekeys ...string,
) ([]notemodel.Note, error) {
	var result []notemodel.Note
	db := s.db

	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}

	db = db.Table(notemodel.Note{}.TableName()).Where(conditions)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
