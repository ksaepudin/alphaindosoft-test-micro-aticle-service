package helper

import (
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	"github.com/labstack/echo/v4"
)

func Response(c echo.Context, errorcode int, message string, data ...interface{}) error {
	return c.JSON(errorcode, &entity.Responses{
		Res: entity.Response{
			Code:    errorcode,
			Message: message,
			Data:    data,
		},
	})
}
