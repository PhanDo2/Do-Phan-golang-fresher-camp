package userbiz

import (
	"Fresher_go/component"
	"Fresher_go/modules/user/usermodel"
	"context"
)

type loginDb interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}
type loginBusiness struct {
	appCtx    component.AppContext
	storeUser loginDb
	hasher    Hasher
}
