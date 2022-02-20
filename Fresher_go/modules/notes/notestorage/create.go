package notestorage

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
)

func (s *sqlStore) Create(cxt context.Context, data *notemodel.NoteCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
