package categoryrepo

import (
	"fmt"
	"projectStaff/model"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

// NewUserRepo ..

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

// GetExistingCategory fetches a category by the id from the db and returns it.

func (repo *CategoryRepo) GetExistingCategory(id int) model.Category {

	var category model.Category // creamos una variable de objeto de categoria.
	repo.db.Where("id = ?", id).First(&category)
	return category
}

// CreateCategory creates a new category in the db..
func (repo *CategoryRepo) CreateCategory(category model.Category) (model.Category, error) {

	result := repo.db.Create(&category)
	fmt.Println(result)
	// result := h.db.Create(&user)
	//  if result 

	fmt.Println("Inserted a user with ID:", category.ID)
	return category, nil
}
