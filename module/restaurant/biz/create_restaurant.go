package restaurantbiz

import (
	restaurantmodel "awesomeProject/module/restaurant/model"
	"context"
)

// khai báo interface ở nơi dùng nó.
type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

// thực thi logic và validate ở đây rồi sẽ đẩy vào db
func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := biz.store.CreateRestaurant(context, data); err != nil {
		return err
	}
	return nil
}
