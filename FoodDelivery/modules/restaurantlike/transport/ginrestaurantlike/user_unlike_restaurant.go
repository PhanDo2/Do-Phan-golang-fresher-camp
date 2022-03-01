package ginrestaurantlike

import (
	"FoodDelivery/common"
	"FoodDelivery/component"
	rstlikebiz "FoodDelivery/modules/restaurantlike/biz"
	restaurantlikestorage "FoodDelivery/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserUnlikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		//data := restaurantlikemodel.Like{
		//	RestaurantId: int(uid.GetLocalID()),
		//	UserId:       requester.GetUserId(),
		//}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserUnlikeRestaurantBiz(store, appCtx.GetPubsub())

		if err := biz.UnlikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
