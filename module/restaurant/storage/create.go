package restaurantstorage

import (
	restaurantmodel "awesomeProject/module/restaurant/model"
	"context"
)

// hàm này sẽ được gắn vào struct sqlStore
// hàm liên quan tới Io nên có context để truy gốc được
func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
