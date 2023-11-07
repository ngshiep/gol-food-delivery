package restaurantstorage

import (
	restaurantmodel "awesomeProject/module/restaurant/model"
	"context"
)

func (s *sqlStore) FindRestaurantWithCondition(
	context context.Context, condition map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {

	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
