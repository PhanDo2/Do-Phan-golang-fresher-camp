package ginrestaurant

import (
	"FoodDelivery/common"
	"FoodDelivery/component"
	"FoodDelivery/modules/restaurant/restaurantbiz"
	"FoodDelivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.ErrInvalidRequest(err)
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		data, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
