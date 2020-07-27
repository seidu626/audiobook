package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

/**
 * GORM
**/

// BeforeCreate implements the GORM BeforeCreate interface for the UserORM type.
// you can use this method to generate new UUID for CREATE operation or let database create it with this annotation:
// {type: "uuid", primary_key: true, not_null:true, default: "uuid_generate_v4()"}];
// we prefer First method as it works with both SQLite & PostgreSQL
func (m *Language) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}

// BeforeCreate implements the GORM BeforeCreate interface for the Skill type.
func (m *Skill) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}

// BeforeCreate implements the GORM BeforeCreate interface for the Word type.
func (m *Word) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}
