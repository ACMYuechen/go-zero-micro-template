// Code scaffolded by goctl. Safe to edit.
// gorm

package file

import "gorm.io/gorm"

var _ FileModel = (*customFileModel)(nil)

type (
	// FileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileModel.
	FileModel interface {
		fileModel
	}

	customFileModel struct {
		*defaultFileModel
	}
)

// NewFileModel returns a model for the database table.
func NewFileModel(conn *gorm.DB) FileModel {
	return &customFileModel{
		defaultFileModel: newFileModel(conn),
	}
}
