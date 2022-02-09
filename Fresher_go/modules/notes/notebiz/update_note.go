package notebiz

//
//import (
//	"Fresher_go/modules/notes/notemodel"
//	"context"
//	"errors"
//)
//
//// khai báo dưới dạng interface
//type UpdateNoteStore interface {
//	//hàm tìm dữ liệu
//	FindDataByCondition(
//		cxt context.Context,
//		conditions map[string]interface{},
//		morekeys ...string,
//	) ([]notemodel.Note, error)
//	//hàm update
//	UpdateData(
//		cxt context.Context,
//		id int, //id cần update
//		data *notemodel.NoteUpdate, // dữ liệu cần update
//	) error
//}
//
//type updateNoteBiz struct {
//	store UpdateNoteStore
//}
//
//func NewUpdateNoteBiz(store UpdateNoteStore) *updateNoteBiz {
//	return &updateNoteBiz{store: store}
//}
//
////step 1: find
//func (biz *updateNoteBiz) UpdateNote(cxt context.Context, id int) error {
//	data, err := biz.store.FindDataByCondition(cxt, map[string]interface{}{"id": id})
//
//	if err != nil {
//		return errors.New("you can't edit")
//	}
//
//	return nil
//}
