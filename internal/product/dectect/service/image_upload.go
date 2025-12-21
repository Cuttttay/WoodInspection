package service

import (
	"WoodInspection/internal/product/dectect/repository"
	"context"
	"time"
)

type ImageUploadService interface {
	SaveImageToServer(ctx context.Context, filePath string) (*repository.Image, error)
}

type ImageUploadServiceImpl struct {
	imageRepo repository.ImageRepository
}

func NewImageUploadService() ImageUploadService {
	return &ImageUploadServiceImpl{imageRepo: imageRepo}
}

func (s *ImageUploadServiceImpl) SaveImageToServer(ctx context.Context, filePath string) (*repository.Image, error) {
	image := &repository.Image{
		FilePath: filePath,
		Uploaded: time.Now(),
	}
	err := s.imageRepo.SaveImage(ctx, image)
	if err != nil {
		return nil, err
	}
	return image, nil
}
