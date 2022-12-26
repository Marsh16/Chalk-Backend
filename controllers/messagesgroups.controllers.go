package controllers

import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)

//add, fetch 
func StoreMessagesGroups(c echo.Context) error{
	messages:= c.FormValue("messages")
	group_id:= c.FormValue("group_id")


	
		result, err := models.StoreMessagesGroups(messages,group_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}

		func FetchMessagesByGroupId(c echo.Context) error{
			group_id:=  c.FormValue("group_id")
		
			
		
				result, err := models.FetchMessagesByGroupId(group_id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}