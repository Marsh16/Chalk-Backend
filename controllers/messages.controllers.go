package controllers

import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)

func StoreMessages(c echo.Context) error{
	messages:= c.FormValue("messages")
	contact_id:= c.FormValue("contact_id")


	
		result, err := models.StoreMessages(messages,contact_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}

		func FetchMessagesByContactId(c echo.Context) error{
			contact_id:=  c.FormValue("contact_id")
		
			
		
				result, err := models.FetchMessagesByContactId(contact_id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}