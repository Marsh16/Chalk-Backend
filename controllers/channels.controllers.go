package controllers


import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)

func FetchAllChannels(c echo.Context) error{
	user_id:=  c.FormValue("user_id")
	

		result, err := models.FetchAllChannels(user_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}