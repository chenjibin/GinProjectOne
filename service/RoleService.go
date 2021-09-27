package service

import (
	"GinProjectOne/common"
	"GinProjectOne/model"
	"gorm.io/gorm"
)

type IRoleService interface {
	Create(name string) (*model.Role, error)
	Update(role model.Role, name string) (*model.Role, error)
	SelectById(id int) (*model.Role, error)
	DeleteById(id int) error
}

type RoleService struct {
	DB *gorm.DB
}

func (c RoleService) DeleteById(id int) error {
	if err := c.DB.Delete(model.Role{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (c RoleService) SelectById(id int) (*model.Role, error) {
	var role model.Role
	if err := c.DB.First(&role, id).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (c RoleService) Create(name string) (*model.Role, error) {
	role := model.Role{
		Name: name,
	}
	if err := c.DB.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (c RoleService) Update(role model.Role, name string) (*model.Role, error) {
	if err := c.DB.Model(&role).Update("name", name).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func NewRoleService() IRoleService {
	return RoleService{DB: common.GetDB()}
}
