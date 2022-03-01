package ginrestaurantlike

import (
	"FoodDelivery/common"
	"FoodDelivery/component"
	rstlikebiz "FoodDelivery/modules/restaurantlike/biz"
	restaurantlikemodel "FoodDelivery/modules/restaurantlike/model"
	restaurantlikestorage "FoodDelivery/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//var filter restaurantlikemodel.Filter
		//
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
