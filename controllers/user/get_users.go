package user

import (
	"net/http"
	"unjuk_prakerja/configs"
	basemodel "unjuk_prakerja/models/base"
	usermodel "unjuk_prakerja/models/users"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []usermodel.User
	result := configs.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, basemodel.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, basemodel.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    users,
	})
}
