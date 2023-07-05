// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSetting = "setting"

// Setting mapped from table <setting>
type Setting struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    string         `gorm:"column:user_id;not null" json:"user_id"` // 用户id
	TeamID    string         `gorm:"column:team_id;not null" json:"team_id"` // 当前团队id
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Setting's table name
func (*Setting) TableName() string {
	return TableNameSetting
}
