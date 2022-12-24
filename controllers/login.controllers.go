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
// func Logout(c echo.Context) error{

// }

func CheckLogin(c echo.Context) error{
	// var ress models.Response
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	//
	res, err := models.CheckLogin(username,password)
	// result, err := models.FetchUsersByUsername(username)
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
	// ress.Status = http.StatusOK
	// ress.Message = "Login Succesfull"
	// ress.Data = result
	return c.JSON(http.StatusOK, 
	map[string]string{
"message": "Success",
"token": mytoken,
"username":username})
// return c.JSON(http.StatusOK, result)
// return c.JSON(http.StatusOK, result)

}
func CheckUserExist(c echo.Context) error{
	// var ress models.Response
	username := c.FormValue("username")
	
	//
	res, err := models.CheckUserExist(username)
	// result, err := models.FetchUsersByUsername(username)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"message":err.Error(),
		})
	}
if !res{
	return echo.ErrUnauthorized
}
	return c.JSON(http.StatusOK, 
	map[string]string{
"message": "Success"})
// return c.JSON(http.StatusOK, result)
// return c.JSON(http.StatusOK, result)

	}
	