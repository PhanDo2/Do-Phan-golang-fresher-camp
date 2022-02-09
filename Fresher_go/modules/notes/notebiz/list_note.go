package notebiz

import (
	"Fresher_go/common"
	"Fresher_go/modules/notes/notemodel"
	"context"
)

type ListNoteStore interface {
	ListDataByCondition(cxt context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		morekeys ...string,
	) ([]notemodel.Note, error)
}
type listNoteBiz struct {
	store ListNoteStore
}

func NewListNoteBiz(store ListNoteStore) *listNoteBiz {
	return &listNoteBiz{store: store}
}

func (biz *listNoteBiz) ListNote(cxt context.Context, paging *common.Paging) ([]notemodel.Note, error) {
	result, err := biz.store.ListDataByCondition(cxt, nil, paging)
	return result, err
}
