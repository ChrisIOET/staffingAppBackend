package userskillrepo

import (
	"fmt"
	"projectStaff/model"
	"gorm.io/gorm"
)

type UserSkillRepo struct {
	db *gorm.DB
}

// NewUserRepo ..
func NewUserSkillRepo(db *gorm.DB) *UserSkillRepo {
	return &UserSkillRepo{
		db: db,
	}
}

// GetExistingUserSkill fetches a skill by the id from the db and returns it.
func (repo *UserSkillRepo) GetExistingUserSkill(id int) model.UserSkill {

	var userSkill model.UserSkill // creamos una variable de objeto de categoria.
	repo.db.Where("id = ?", id).First(&userSkill)
	return userSkill
}

// CreateUserSkill creates a new userSkill in the db..
func (repo *UserSkillRepo) CreateUserSkill(userSkill model.UserSkill) (model.UserSkill, error) {

	result := repo.db.Create(&userSkill)
	fmt.Println(result)
	// result := h.db.Create(&user)
	//  if result 

	fmt.Println("Inserted a user with ID:", userSkill.UserID)
	return userSkill, nil
}
