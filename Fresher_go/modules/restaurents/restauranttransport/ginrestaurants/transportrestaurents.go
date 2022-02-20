package ginrestaurants

import (
	"Fresher_go/common"
	"Fresher_go/component"
	restaurantlikestorage "Fresher_go/modules/restaurentlike/storage"
	"Fresher_go/modules/restaurents/restaurantbiz"
	"Fresher_go/modules/restaurents/restaurantdb"
	"Fresher_go/modules/restaurents/restaurantmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRetaurents(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.Restaurants
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantdb.NewSQLStore(AppCtx.GetMainDBConnection())
		biz := restaurantbiz.NewRestaurants(store)
		if err := biz.CreateRestaurants(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSONP(http.StatusOK, common.SimpleSuccesResponse(data.FakeId.String()))
	}

}
func ListRetaurents(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var paging common.Paging
		//if err := c.ShouldBind(&paging); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//paging.Fulfill()
		//
		//store := restaurantdb.NewSQLStore(AppCtx.GetMainDBConnection())
		//biz := restaurantbiz.NewRestaurants(store)
		//data, err := biz.ListRestaurants(c.Request.Context(), &paging)
		//if err != nil {
		//	panic(err)
		//}
		//for i := range data {
		//	data[i].Maks(false)
		//	if i == len(data)-1 {
		//		paging.NextCursor = data[i].FakeId.String()
		//	}
		//}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantdb.NewSQLStore(AppCtx.GetMainDBConnection())
		likeStore := restaurantlikestorage.NewSQLStore(AppCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store, likeStore)

		result, err := biz.ListRestaurants(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}
		c.JSONP(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}

}
func UpdateRetaurents(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.Restaurants
		if err := c.ShouldBind(&data); err != nil {
			// change err message and err code
			c.JSONP(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		store := restaurantdb.NewSQLStore(AppCtx.GetMainDBConnection())
		biz := restaurantbiz.NewRestaurants(store)
		if err := biz.UpdateRestaurants(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			c.JSONP(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSONP(http.StatusOK, common.SimpleSuccesResponse(true))
	}

}
func DeleteRetaurents(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Param("id") lấy id truyền vào từ param
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := restaurantdb.NewSQLStore(AppCtx.GetMainDBConnection())
		biz := restaurantbiz.NewRestaurants(db)
		if err := biz.DeleteRestaurants(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSONP(http.StatusOK, common.SimpleSuccesResponse(true))
	}

}
func GetRetaurents(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//id, err := strconv.Atoi(c.Param("id"))
		// lỗi từ thư viện chuẩn của go trả về nên cần bọc lại
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantdb.NewSQLStore(AppCtx.GetMainDBConnection())
		biz := restaurantbiz.NewRestaurants(store)
		data, err := biz.FindRestaurants(c.Request.Context(), int(uid.GetLocalID()))
		// lỗi từ BIz trả về nên cần bọc lại vì là AppError sẵn r
		if err != nil {

			panic(err)
		}
		data.Mask(false)
		c.JSONP(http.StatusOK, common.SimpleSuccesResponse(data))
	}

}
