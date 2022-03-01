package ginuser

import (
	"FoodDelivery/common"
	"FoodDelivery/component"
	"FoodDelivery/component/hasher"
	"FoodDelivery/modules/user/userbiz"
	"FoodDelivery/modules/user/usermodel"
	"FoodDelivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
