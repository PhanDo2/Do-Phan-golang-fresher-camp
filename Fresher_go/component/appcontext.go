//package component
//
//import "gorm.io/gorm"
//
//type AppContext interface {
//	GetMainDBConnection() *gorm.DB
//}
//type appCtx struct {
//	db *gorm.DB
//}
//
//func NewAppConntect(db *gorm.DB) *appCtx {
//	return &appCtx{db: db}
//}
//func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
//	return ctx.db
//}
package component

import (
	"Fresher_go/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
