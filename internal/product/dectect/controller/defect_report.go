// controller/defect_report.go
package controller

import (
	"WoodInspection/internal/product/dectect/repository"
	"github.com/gin-gonic/gin"
)

func GetDefectReport(c *gin.Context) {
	reportID := c.Param("id")

	// 获取检测结果
	defectRepo := repository.NewDefectRepository()
	defects, err := defectRepo.GetDefectsByImageID(c, reportID)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取报告失败"})
		return
	}

	c.JSON(200, gin.H{
		"defects": defects,
	})
}
