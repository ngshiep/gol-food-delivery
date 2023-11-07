package restaurantbiz

import (
	restaurantmodel "awesomeProject/module/restaurant/model"
	"context"
)

// khai báo interface ở nơi dùng nó.
type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

// thực thi logic và validate ở đây rồi sẽ đẩy vào db
// implement func của interface
func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.Create(context, data); err != nil {
		return err
	}
	return nil
}
