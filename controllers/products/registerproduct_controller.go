package prodcuts

import (
	"net/http"
	"unjuk_prakerja/configs"
	"unjuk_prakerja/models/base"
	productsmodel "unjuk_prakerja/models/product"
	usermodels "unjuk_prakerja/models/users"

	"github.com/labstack/echo/v4"
)

func RegisterProduct(c echo.Context) error {
	var product productsmodel.ProductData
	var user usermodels.User
	c.Bind(&product)

	checkUser := configs.DB.Where("name=?", product.Publisher).First(&user)
	if checkUser.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Wrong Publisher Name/ Publisher have not registered yet",
			Data:    nil,
		})
	}

	addProduct := configs.DB.Create(&product)
	if addProduct.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed adding data to database!",
			Data:    nil,
		})
	}

	var authResponse = productsmodel.ProductData{
		ID:        product.ID,
		Name:      product.Name,
		Stock:     product.Stock,
		Price:     product.Price,
		Publisher: product.Publisher,
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Successfully Adding Data!",
		Data:    authResponse,
	})

}
