// defect_repository.go - 存储缺陷检测结果
package repository

import (
	"WoodInspection/internal/product/dectect/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DefectReaderRepository interface {
	GetDefectsByImageID(c *gin.Context, ImageID string) ([]model.Defect, error)
}

type DefectWriterRepository interface {
	SaveDefect(c *gin.Context, defect *model.Defect) error
}

type DefectRepository struct {
	defectReader DefectReaderRepository
	defectWriter DefectWriterRepository
}

type DefectGormRepository struct {
	gormDB *gorm.DB
}

func NewDefectRepository() *DefectRepository {
	return &DefectRepository{}
}

func (u DefectGormRepository) GetDefectsByImageID(c *gin.Context, ImageID string) ([]model.Defect, error) {
	defect := model.Defect{}
	err := u.gormDB.Where("image_id = ?", ImageID).First(&defect).Error
	if err != nil {
		return []model.Defect{}, err
	}
	return nil, err
}

func (u *DefectGormRepository) SaveDefect(c *gin.Context, defect model.Defect) error {
	defect = model.Defect{}
	err := u.gormDB.Create(&defect).Error
	if err != nil {
		return err
	}
	return nil
}
