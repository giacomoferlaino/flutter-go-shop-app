package orm

import "time"

// Model is the base orm model containing mandatory fields
type Model struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
