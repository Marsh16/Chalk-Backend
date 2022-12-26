package controllers

import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)

//add, fetch 
func StoreContactsGroups(c echo.Context) error{
	contact_id:= c.FormValue("contact_id")
	group_id:= c.FormValue("group_id")


	
		result, err := models.StoreContactsGroups(contact_id,group_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}

		func FetchContactByGroupId(c echo.Context) error{
			group_id:=  c.FormValue("group_id")
		
			
		
				result, err := models.FetchContactByGroupId(group_id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}