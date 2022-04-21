package repository

import (
	"projectStaff/repository/categoryrepo"
	"projectStaff/repository/skillrepo"
	"projectStaff/repository/userrepo"
	"projectStaff/repository/userskillrepo"
	"gorm.io/gorm"
)

// Repositories contains all the repo structs
type Repositories struct {
	UserRepo     *userrepo.UserRepo 		//UserRepo = Structure of UserRepo
	CategoryRepo *categoryrepo.CategoryRepo
	SkillRepo    *skillrepo.SkillRepo
	UseSkillRepo *userskillrepo.UserSkillRepo
}

// InitRepositories should be called in main.go

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := userrepo.NewUserRepo(db)
	categoryRepo := categoryrepo.NewCategoryRepo(db)
	skillrepo := skillrepo.NewSkillRepo(db)
	userskillrepo := userskillrepo.NewUserSkillRepo(db)

	return &Repositories{
		UserRepo:     userRepo,
		CategoryRepo: categoryRepo,
		SkillRepo:    skillrepo,
		UseSkillRepo: userskillrepo,
	}
}
