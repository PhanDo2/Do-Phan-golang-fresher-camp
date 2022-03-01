package restaurantbiz

import (
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"context"
)

type CreateRestaurantStore interface {
	CreateData(
		ctx context.Context,
		data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}
func (biz *createRestaurantBiz) CreateRestaurant(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	if err := biz.store.CreateData(ctx, data); err != nil {
		return err
	}
	return nil
}
