package restaurantbiz

import (
	restaurantmodel "awesomeProject/module/restaurant/model"
	"context"
	"errors"
)

// khai báo interface ở nơi dùng nó.
type DeleteRestaurantStore interface {
	FindRestaurantWithCondition(
		context context.Context, condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)

	Delete(context context.Context, id int) error
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {

	oldData, err := biz.store.FindRestaurantWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := biz.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
