// Code scaffolded by goctl. Safe to edit.
// gorm

package file

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type (
	fileModel interface {
		CreateTable() error
		Insert(ctx context.Context, data []*File) error
		InsertOne(ctx context.Context, data *File) error
		FindOne(ctx context.Context, id int64) (*File, error)
		Update(ctx context.Context, data *File) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFileModel struct {
		conn  *gorm.DB
		table string
	}

	File struct {
		Id           int64          `gorm:"primaryKey;autoIncrement;comment:主键ID"`
		FileName     string         `gorm:"size:255;comment:文件名"`
		FilePath     string         `gorm:"size:500;comment:文件路径"`
		FileSize     int64          `gorm:"comment:文件大小"`
		FileType     string         `gorm:"size:100;comment:文件类型"`
		UploadUserID string         `gorm:"comment:上传用户ID"`
		IsPrivate    bool           `gorm:"comment:是否私有"`
		CreatedAt    time.Time      `gorm:"autoCreateTime;comment:创建时间"`
		UpdatedAt    time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
		DeletedAt    gorm.DeletedAt `gorm:"index;comment:删除时间"`
	}
)

func newFileModel(conn *gorm.DB) *defaultFileModel {
	return &defaultFileModel{
		conn:  conn,
		table: "`file`",
	}
}
func (m *defaultFileModel) Delete(ctx context.Context, id int64) error {
	return m.conn.WithContext(ctx).Where("id = ?", id).Delete(&File{}).Error
}

func (m *defaultFileModel) FindOne(ctx context.Context, id int64) (*File, error) {
	model := &File{}
	err := m.conn.WithContext(ctx).Where("id = ?", id).First(model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultFileModel) CreateTable() error {
	return m.conn.AutoMigrate(&File{})
}

func (m *defaultFileModel) Insert(ctx context.Context, data []*File) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *defaultFileModel) InsertOne(ctx context.Context, data *File) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *defaultFileModel) Update(ctx context.Context, data *File) error {
	return m.conn.WithContext(ctx).Save(data).Error
}

func (m *defaultFileModel) tableName() string {
	return m.table
}
