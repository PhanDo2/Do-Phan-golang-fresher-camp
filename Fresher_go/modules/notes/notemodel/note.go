package notemodel

import (
	"errors"
	"strings"
)

// tại đây thì lưu dữ liệu trả về tương tự như model trong MVC
type Note struct {
	Id     int    `json:"id,omitempty" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Addr   string `json:"addr" gorm:"column:addr"`
	Status int    `json:"status" gorm:"column:status"`
}

// tại đây thì lưu dữ liệu trả về tương tự như model trong MVC
func (Note) TableName() string {
	return "note"
}

type NoteUpdate struct {
	Name   *string `json:"name" gorm:"column:name"`
	Addr   *string `json:"addr" gorm:"column:addr"`
	Status *int    `json:"status" gorm:"column:status"`
}

func (NoteUpdate) TableName() string {
	return Note{}.TableName()
}

type NoteCreate struct {
	Id     int    `json:"id,omitempty" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Addr   string `json:"addr" gorm:"column:addr"`
	Status int    `json:"status" gorm:"column:status"`
}

func (NoteCreate) TableName() string {
	return Note{}.TableName()
}

func (note *NoteCreate) Validate() error {
	note.Name = strings.TrimSpace(note.Name)
	if len(note.Name) == 0 {
		return errors.New("note name can not be blank")
	}
	return nil

}
