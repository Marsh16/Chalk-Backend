package controllers

import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)


func FetchAllUsers(c echo.Context) error{
	result, err := models.FetchAllUsers()
	
	if err != nil{
		return c.JSON(http.StatusInternalServerError,
		result)
	}
	return c.JSON(http.StatusOK, result)
	}

func StoreUsers(c echo.Context) error{
name:= c.FormValue("name")
username:= c.FormValue("username")
email:= c.FormValue("email")
phone_number:= c.FormValue("phone_number")
dateofbirth:= c.FormValue("dateofbirth")
profilepic:= c.FormValue("profilepic")
password:=  c.FormValue("password")
	result, err := models.StoreUsers(name,username,email,phone_number,dateofbirth, profilepic,password)
	
	if err != nil{
		return c.JSON(http.StatusInternalServerError,
		map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
	}

	func UpdateUsers(c echo.Context) error{
		id:=  c.FormValue("user_id")
		name:= c.FormValue("name")
username:= c.FormValue("username")
email:= c.FormValue("email")
phone_number:= c.FormValue("phone_number")
dateofbirth:= c.FormValue("dateofbirth")
profilepic:= c.FormValue("profilepic")
password:=  c.FormValue("password")

			result, err := models.UpdateUsers(id,name,username,email,phone_number,dateofbirth,profilepic,password)
			
			if err != nil{
				return c.JSON(http.StatusInternalServerError,
				map[string]string{"message": err.Error()})
			}
			return c.JSON(http.StatusOK, result)
			}

			func DeleteUsers(c echo.Context) error{
				id:= c.FormValue("user_id")
				
					result, err := models.DeleteUsers(id)
					
					if err != nil{
						return c.JSON(http.StatusInternalServerError,
						map[string]string{"message": err.Error()})
					}
					return c.JSON(http.StatusOK, result)
					}