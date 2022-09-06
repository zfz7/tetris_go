package server

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zfz7/tetris_go/backend/repository"
	"gorm.io/gorm"
	"net/http"
)

type HelloDTO struct {
	Message string `json:"message" xml:"message"`
}

func helloWorldEndPoint(context echo.Context) error {
	var metaData repository.MetaData
	db, _ := repository.GetDB()
	if err := db.First(&metaData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			metaData = repository.MetaData{UserCount: 0}
			db.Create(&metaData)
		} else {
			panic(err)
		}
	}
	dto := &HelloDTO{
		Message: fmt.Sprint("You are visitor ", metaData.UserCount),
	}
	db.Model(&metaData).Update("UserCount", metaData.UserCount+1)
	return context.JSON(http.StatusOK, dto)
}
