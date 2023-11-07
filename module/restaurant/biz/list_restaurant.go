package restaurantbiz

import (
	"awesomeProject/common"
	restaurantmodel "awesomeProject/module/restaurant/model"
	"context"
)

// khai báo interface ở nơi dùng nó.
type ListRestaurantStore interface {
	ListRestaurantWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

// thực thi logic và validate ở đây rồi sẽ đẩy vào db
// implement func của interface
func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListRestaurantWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
