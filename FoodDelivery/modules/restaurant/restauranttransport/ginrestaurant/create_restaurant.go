package ginrestaurant

import (
	"FoodDelivery/common"
	"FoodDelivery/component"
	"FoodDelivery/modules/restaurant/restaurantbiz"
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"FoodDelivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			common.ErrInvalidRequest(err)
		}

		res := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = res.GetUserId()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.GenUID(common.DbTypeRestaurant)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
