package notestorage

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
)

func (s *sqlStore) DeleteData(
	cxt context.Context,
	id int,
	data *notemodel.Note,
) error {
	db := s.db
	if err := db.Where("id= ?", id).Delete(data).Error; err != nil {

		return err
	}
	return nil
}
