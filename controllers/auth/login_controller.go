package auth

import (
	"net/http"
	"unjuk_prakerja/configs"
	"unjuk_prakerja/models/base"
	usermodel "unjuk_prakerja/models/users"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var user usermodel.User
	c.Bind(&user)

	checkuser := configs.DB.Where("name=?", user.Name).First(&user)
	if checkuser.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Wrong Username",
		})
	}

	var authResponse = usermodel.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success Loging in!, Here's your data :",
		Data:    authResponse,
	})
}
