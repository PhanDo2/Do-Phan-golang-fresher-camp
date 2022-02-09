package notestorage

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
)

func (s *sqlStore) UpdateData(
	cxt context.Context,
	//id cần update
	id int,
	// dữ liệu cần update
	data *notemodel.NoteUpdate,
) error {
	db := s.db
	if err := db.Where("id= ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
