package rstlikebiz

import (
	"FoodDelivery/common"
	restaurantlikemodel "FoodDelivery/modules/restaurantlike/model"
	"FoodDelivery/pubsub"
	"context"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

//
//type DecreaseLikeCountStore interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
	//decStore DecreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
	//decStore DecreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
		//decStore: decStore,
		pubsub: pubsub,
	}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{
		RestaurantId: restaurantId,
		UserId:       userId,
	}))

	return nil
}
