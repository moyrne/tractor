package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		return 1
	}

	return page
}

const (
	minPageSize = 10
	maxPageSize = 100
)

func GetPageSize(c *gin.Context) int {
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		return minPageSize
	}
	if pageSize > maxPageSize {
		return maxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
