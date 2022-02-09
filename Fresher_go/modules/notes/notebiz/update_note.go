package notebiz

import (
	"Fresher_go/modules/notes/notemodel"
	"context"
	"errors"
)

// khai báo dưới dạng interface
type UpdateNoteStore interface {
	//hàm tìm dữ liệu
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		morekeys ...string,
	) (*notemodel.Note, error)

	//FindDataByUpdate(
	//	ctx context.Context,
	//	conditions map[string]interface{},
	//	morekeys ...string,
	//) (notemodel.Note, error)

	// hàm update
	UpdateData(
		cxt context.Context,
		//id cần update
		id int,
		// dữ liệu cần update
		data *notemodel.NoteUpdate,
	) error
}

type updateNoteBiz struct {
	store UpdateNoteStore
}

func NewUpdateNoteBiz(store UpdateNoteStore) *updateNoteBiz {
	return &updateNoteBiz{store: store}
}

//step 1: find
func (biz *updateNoteBiz) UpdateNote(ctx context.Context, id int, data *notemodel.NoteUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if oldData.Id == 0 {
		return errors.New("data not found")
	}
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
