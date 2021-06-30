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

type Map struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (m *Map) Save(tx *gorm.DB) error {
	return tx.Create(m).Error
}

func (m *Map) Load(tx *gorm.DB, id uint) error {
	return tx.Take(m, id).Error
}

func (m *Map) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(m, id).Error
}

type Maps []Map

func (m *Maps) Load(tx *gorm.DB) error {
	return tx.Find(m).Order("ID asc").Error
}

type Tip struct {
	gorm.Model
	Title            string `gorm:"not null"`
	StandingPosition string `gorm:"column:standing_position;not null"`
	AimPosition      string `gorm:"column:aim_position;not null"`
	Description      string `gorm:"not null"`
	SideID           uint   `gorm:"not null"`
	Side             Side
	MapID            uint `gorm:"not null"`
	Map              Map
	AbilityID        uint `gorm:"not null"`
	Ability          Ability
	Good             int `gorm:"not null;default:0"`
	Bad              int `gorm:"not null;default:0"`
}

type Side struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"not null"`
}

func (s *Side) Save(tx *gorm.DB) error {
	return tx.Create(s).Error
}

func (t *Tip) Save(tx *gorm.DB) error {
	return tx.Create(t).Error
}

func (t *Tip) Load(tx *gorm.DB, id uint) error {
	return tx.Take(t, id).Error
}

func (t *Tip) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(t, id).Error
}

type Sides []Side

func (s *Sides) Load(tx *gorm.DB) error {
	return tx.Find(s).Order("ID asc").Error
}

type Tips []Tip

func (t *Tips) Load(tx *gorm.DB) error {
	return tx.Preload("Side").Preload("Map").Preload("Ability").Find(t).Order("ID asc").Error
}
