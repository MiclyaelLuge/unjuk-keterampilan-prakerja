package auth

import (
	"net/http"
	"unjuk_prakerja/configs"
	"unjuk_prakerja/middleware"
	"unjuk_prakerja/models/base"
	usermodel "unjuk_prakerja/models/users"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var user usermodel.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := configs.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Faild adding data to database",
			Data:    nil,
		})
	}

	var authResponse = usermodel.ResponseAuth{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.ID, user.Name),
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Successfully Adding Data!",
		Data:    authResponse,
	})

}
