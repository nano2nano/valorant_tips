package model

import (
	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	Name     string `gorm:"column:name;not null"`
	Ability1 int    `gorm:"column:ability1;not null"`
	Ability2 int    `gorm:"column:ability2;not null"`
	Ability3 int    `gorm:"column:ability3;not null"`
	Ability4 int    `gorm:"column:ability4;not null"`
}

func (i *Agent) Save(tx *gorm.DB) error {
	return tx.Create(i).Error
}

func (i *Agent) Load(tx *gorm.DB, id uint) error {
	return tx.Take(i, id).Error
}

func (i *Agent) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(i, id).Error
}

type Inputs []Agent

func (i *Inputs) Load(tx *gorm.DB) error {
	return tx.Find(i).Order("ID asc").Error
}
