package restaurantbiz

import (
	"Fresher_go/common"
	"Fresher_go/modules/restaurents/restaurantmodel"
	"context"
	"log"
)

type RestaurentStore interface {
	Create(ctx context.Context, data *restaurantmodel.Restaurants) error
	Find(cxt context.Context, conditions map[string]interface{}, morekeys ...string) (*restaurantmodel.Restaurants, error)
	Update(cxt context.Context, id int, data *restaurantmodel.Restaurants) error
	Delete(
		ctx context.Context,
		id int,
	) error
	List(cxt context.Context, conditions map[string]interface{}, paging *common.Paging, morekeys ...string) ([]restaurantmodel.Restaurants, error)
}
type RestaurentBiz struct {
	store RestaurentStore
}

func NewRestaurants(store RestaurentStore) *RestaurentBiz {
	return &RestaurentBiz{store: store}
}

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	store     RestaurentStore
	likeStore LikeStore
}

func NewListRestaurantBiz(store RestaurentStore, likeStore LikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

// Func logic
func (biz *RestaurentBiz) CreateRestaurants(ctx context.Context, data *restaurantmodel.Restaurants) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)
	return err

}

//update
func (biz *RestaurentBiz) UpdateRestaurants(ctx context.Context, id int, data *restaurantmodel.Restaurants) error {
	// check input
	oldData, err := biz.store.Find(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}
	if oldData.Id == 0 {
		return common.RecordNotFound
	}
	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}

//softdelete
func (biz *RestaurentBiz) DeleteRestaurants(ctx context.Context, id int) error {
	// check input
	oldData, err := biz.store.Find(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}
	if oldData.Id == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}
	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}

//list data
func (biz *listRestaurantBiz) ListRestaurants(ctx context.Context, paging *common.Paging) ([]restaurantmodel.Restaurants, error) {
	// check input
	result, err := biz.store.List(ctx, nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println("Cannot get restaurant likes:", err)
	}

	if v := mapResLike; v != nil {
		for i, item := range result {
			result[i].LikedCount = mapResLike[item.Id]
		}
	}

	return result, nil
}

//find
func (biz *RestaurentBiz) FindRestaurants(ctx context.Context, id int) (*restaurantmodel.Restaurants, error) {
	//data, err := biz.store.Find(ctx, map[string]interface{}{"id": id})
	//if err != nil {
	//
	//	if err != common.RecordNotFound {
	//		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	//
	//	}
	//	return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	//}
	//if data.Status == 0 {
	//	return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	//}
	//return data, err
	data, err := biz.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return data, nil
}
