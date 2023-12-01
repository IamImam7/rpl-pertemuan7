package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	err := c.db.Create(Category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	result := c.db.Model(&model.Category{}).Where("Id = ?",id).Updates(category)
	if result.Error!= nil{
		return result.Error
	}
	return nil // TODO: replace this
}

func (c *categoryRepository) Delete(id int) error {
	result := c.db.Delete(&model.Category{}, id)
	if result.Error !=nil{
		return result.Error
	}
	return nil 
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	var Category model.Category
	err := c.db.Where("id = ?", id).First(&Category).Error
	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	result := []model.Category{}
	rows,err := c.db.Model([]model.Category{}).Select("*").Rows()
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	for rows.Next(){
		var hsl model.Category
		if err := c.db.ScanRows(rows, &hsl); err!=nil{
			return nil,err
		}
		result = append(result,hsl)
	}
	if err:= rows.Err();err!=nil{
		return nil,err
	}
	return result, nil // TODO: replace this
}
