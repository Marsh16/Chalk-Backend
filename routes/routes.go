package routes
// go mod init vp_week11_echo 
// GO111MODULE=on go get github.com/labstack/echo/v4

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"vp_week11_echo/controllers"
	// "vp_week11_echo/middleware"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	name := c.Param("name")
  return c.String(http.StatusOK, "Hello, " + name)
}


func Init() *echo.Echo {
	e := echo.New()
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	
	//untuk user/profile
	e.POST("/login", controllers.CheckLogin)
	e.POST("/loginsuccess", controllers.FetchAllUsers)
	e.POST("/channels", controllers.FetchAllChannels)
	//registration
	e.POST("/users", controllers.StoreUsers)
	//edit profile
	e.PATCH("/users", controllers.UpdateUsers)
	//delete account(mungkin tidak perlu)
	e.DELETE("/users", controllers.DeleteUsers)
	//bisa pengecekan dari app (mungkin tidak perlu validation dari api)
	e.POST("/test-validation", controllers.TestStructValidation)
	e.POST("/test-validation-var", controllers.TestVarValidation)
	//
	return e

}

