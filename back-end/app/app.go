package app

import (
	"github.com/jinzhu/gorm"
)

// State contains the current application state
type State struct {
	Database *gorm.DB
}
