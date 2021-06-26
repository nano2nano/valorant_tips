package model

import (
	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	Name      string    `gorm:"not null"`
	Abilities []Ability `gorm:"not null"`
}

func (a *Agent) Save(tx *gorm.DB) error {
	return tx.Create(a).Error
}

func (a *Agent) Load(tx *gorm.DB, id uint) error {
	return tx.Preload("Abilities").Find(a).Order("ID").Error
}

func (a *Agent) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(a, id).Error
}

type Agents []Agent

func (a *Agents) Load(tx *gorm.DB) error {
	return tx.Preload("Abilities").Find(a).Order("ID asc").Error
}

type Ability struct {
	gorm.Model
	Name        string `gorm:"not null"`
	IconName    string
	Description string `gorm:"not null"`
	AgentID     uint
}

func (a *Ability) Save(tx *gorm.DB) error {
	return tx.Create(a).Error
}

func (a *Ability) Load(tx *gorm.DB, id uint) error {
	return tx.Take(a, id).Error
}

func (a *Ability) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(a, id).Error
}

type Abilities []Ability

func (a *Abilities) Load(tx *gorm.DB) error {
	return tx.Find(a).Order("ID asc").Error
}
