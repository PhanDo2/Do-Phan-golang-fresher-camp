package notebiz

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
	"errors"
)

// khai báo dưới dạng interface
type GetNoteStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*notemodel.Note, error)
}

type getNoteBiz struct {
	store GetNoteStore
}

func NewGetNoteBiz(store GetNoteStore) *getNoteBiz {
	return &getNoteBiz{store: store}
}

func (biz *getNoteBiz) GetNote(cxt context.Context, id int) (*notemodel.Note, error) {
	data, err := biz.store.FindDataByCondition(nil, map[string]interface{}{"id": id})

	if err != nil {
		return nil, errors.New("no")
	}

	return data, err
}
