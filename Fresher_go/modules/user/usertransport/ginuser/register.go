package ginuser

import (
	"Fresher_go/common"
	"Fresher_go/component"
	"Fresher_go/component/hasher"
	"Fresher_go/modules/user/userbiz"
	"Fresher_go/modules/user/userdb"
	"Fresher_go/modules/user/usermodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userdb.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterUserBiz(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccesResponse("Account successfully created"))
	}
}
