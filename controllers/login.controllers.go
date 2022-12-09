package controllers

import(
	"github.com/labstack/echo/v4"
	"vp_week11_echo/helpers"
	"vp_week11_echo/models"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
)
func GenerateHashPassword(c echo.Context) error{
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)

}

func FetchAllUsers(c echo.Context) error{
	result, err := models.FetchAllUsers()
	
	if err != nil{
		return c.JSON(http.StatusInternalServerError,
		result)
	}
	return c.JSON(http.StatusOK, result)
	}
	
func CheckLogin(c echo.Context) error{
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username,password)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"message":err.Error(),
		})
	}
if !res{
	return echo.ErrUnauthorized
}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"]= "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	mytoken, err := token.SignedString([]byte("secret"))
	if err!= nil {
		return c.JSON(http.StatusInternalServerError,map[string]string{"message": err.Error()})

	}

	return c.JSON(http.StatusOK, 
	map[string]string{
"message": "Login Succesfull",
"token": mytoken,
	})

}