package notebiz

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
	"errors"
)

type DeleteNoteStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*notemodel.Note, error)
	DeleteData(
		cxt context.Context,
		id int,
		data *notemodel.Note,
	) error
}

type deleteNoteBiz struct {
	store DeleteNoteStore
}

func NewDeleteNoteBiz(store DeleteNoteStore) *deleteNoteBiz {
	return &deleteNoteBiz{store: store}
}

func (biz *deleteNoteBiz) DeleteNote(ctx context.Context, id int) error {
	data, err := biz.store.FindDataByCondition(nil, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if data.Id == 0 {
		return errors.New("data not found")
	}
	if data.Status == 0 {
		return errors.New("Data has been delete")
	}
	if err := biz.store.DeleteData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
