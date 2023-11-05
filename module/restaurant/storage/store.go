package restaurantstorage

import "gorm.io/gorm"

// s không viết hoa thì k cần phải public ra bên ngoài
type sqlStore struct {
	db *gorm.DB
}

// N viết hoa để public ra bên ngoài giống như hàm contructor còn sqlStore là biến private
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
