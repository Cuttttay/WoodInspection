package repository

import (
	"context"
	"time"
)

// image_repository.go - 用于保存图片信息到数据库
type ImageRepository interface {
	SaveImage(ctx context.Context, image *Image) error
	GetImage(ctx context.Context, id int) (*Image, error)
}

type Image struct {
	ID       int    `json:"id"`
	FilePath string `json:"file_path"`
	Uploaded time.Time
}

func (i *Image) SaveImage(ctx context.Context, image *Image) error {

}

func (i *Image) GetImage(ctx context.Context, id int) (*Image, error) {

}
