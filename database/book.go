package database

import (
	"github.com/balajisainath/restapigo/models"
	"github.com/jinzhu/gorm"
)

// this getbbook func is creating connecton and interaction from golang app to database server via db variable
func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func GetBookByID(db *gorm.DB, id string) (models.Book, error) {

	return models.Book{}, nil
}
func DeleteBookID(db *gorm.DB, id string) error {

	return nil
}
func UpdateBookID(db *gorm.DB, book *models.Book) error {

	return nil
}
