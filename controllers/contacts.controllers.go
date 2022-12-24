package controllers


import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)
func UpdateContacts(c echo.Context) error{
	id:=  c.FormValue("contact_id")
	name:= c.FormValue("name")
phone_number:= c.FormValue("phone_number")
profilepic:= c.FormValue("profilepic")
user_id:= c.FormValue("user_id")

		result, err := models.UpdateContacts(id,name,phone_number,profilepic,user_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func StoreContacts(c echo.Context) error{
	name:= c.FormValue("name")
	phone_number:= c.FormValue("phone_number")
	profilepic:= c.FormValue("profilepic")
	user_id:= c.FormValue("user_id")

	
		result, err := models.StoreContacts(name,phone_number,profilepic,user_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func FetchAllContacts(c echo.Context) error{
	user_id:=  c.FormValue("user_id")
	

		result, err := models.FetchAllContact(user_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func DeleteContacts(c echo.Context) error{
	id:= c.FormValue("contact_id")
			
	result, err := models.DeleteContacts(id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}

func FetchContactsByContactId(c echo.Context) error{
	contact_id:=  c.FormValue("contact_id")

	

		result, err := models.FetchContactsByContactId(contact_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}

		func CheckContactExist(c echo.Context) error{
			// var ress models.Response
			phone_number := c.FormValue("phone_number")
			
			//
			res, err := models.CheckContactExist(phone_number)
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