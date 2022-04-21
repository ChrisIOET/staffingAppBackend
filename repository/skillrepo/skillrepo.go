package skillrepo

import (
	"fmt"
	"projectStaff/model"
	"gorm.io/gorm"
)

type SkillRepo struct {
	db *gorm.DB
}

// NewSKillRepo ..
func NewSkillRepo(db *gorm.DB) *SkillRepo {
	return &SkillRepo{
		db: db,
	}
}

// GetExistingSkill fetches a Skill by the id from the db and returns it.
func (repo *SkillRepo) GetExistingSkill(id int) model.Skill {

	var Skill model.Skill // creamos una variable de objeto de categoria.
	repo.db.Where("id = ?", id).First(&Skill)
	return Skill
}

// CreateSkill creates a new Skill in the db..
func (repo *SkillRepo) CreateSkill(Skill model.Skill) (model.Skill, error) {

	result := repo.db.Create(&Skill)
	fmt.Println(result)
	// result := h.db.Create(&user)
	//  if result 

	fmt.Println("Inserted a user with ID:", Skill.ID)
	return Skill, nil
}
