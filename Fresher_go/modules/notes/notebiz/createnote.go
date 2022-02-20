package notebiz

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
)

type CreateNoteStore interface {
	Create(ctx context.Context, data *notemodel.NoteCreate) error
}
type createNoteBiz struct {
	store CreateNoteStore
}

func NewCreateNoteBiz(store CreateNoteStore) *createNoteBiz {
	return &createNoteBiz{store: store}
}

func (biz *createNoteBiz) Create_Note(ctx context.Context, data *notemodel.NoteCreate) error {
	//if data.Name == "" {
	//	return errors.New("Note name can not be blank")
	//}
	if err := data.Validate(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)

	return err
}
