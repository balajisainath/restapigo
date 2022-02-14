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

func GetBookByID(db *gorm.DB, id string) (*models.Book, error) {
	book:=models.Book{}
	err:=db.Select("Books.*").Group("books.id").Where("books.id=?",id).First(&book).Error
if err!=nil{

	return nil,err
}
return &book,nil
}
func DeleteBookID(db *gorm.DB, id string) error {
var book models.Book
err:=db.Where("id=?",id).Delete(&book).Error
if err!=nil{
	return err
}
	return nil
}
func UpdateBookID(db *gorm.DB, book *models.Book) error {

	err:=db.Save(book).Error
	if err!=nil{
		return err
	}
	return nil
}

func SaveBook(db *gorm.DB, book *models.Book) error {
	err:=db.Save(book).Error
	if err!=nil{
		return err
	}
	return nil

}
