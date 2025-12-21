// controller/image_upload.go
package controller

import (
	"WoodInspection/internal/product/dectect/repository"
	"WoodInspection/internal/product/dectect/service"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "上传图片失败"})
		return
	}

	// 保存图片到服务器
	filePath := "/path/to/save/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(500, gin.H{"error": "保存图片失败"})
		return
	}

	// 调用图片存储服务
	imageService := service.NewImageUploadService()
	image, err := imageService.SaveImageToServer(c, filePath)
	if err != nil {
		c.JSON(500, gin.H{"error": "保存图片信息失败"})
		return
	}

	// 调用缺陷检测服务
	defectService := service.NewDefectDetectionService()
	defects, err := defectService.DetectDefects(image)
	if err != nil {
		c.JSON(500, gin.H{"error": "缺陷检测失败"})
		return
	}

	// 存储缺陷信息
	defectRepo := repository.NewDefectRepository()
	for _, defect := range defects {
		if err := defectRepo.SaveDefect(c, &defect); err != nil {
			c.JSON(500, gin.H{"error": "存储缺陷信息失败"})
			return
		}
	}

	c.JSON(200, gin.H{
		"image":   image,
		"defects": defects,
	})
}
