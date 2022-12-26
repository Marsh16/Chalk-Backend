package controllers

import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)

//add, fetch 
func StoreContactsChannels(c echo.Context) error{
	contact_id:= c.FormValue("contact_id")
	channel_id:= c.FormValue("channel_id")


	
		result, err := models.StoreContactsChannels(contact_id,channel_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}

		func FetchContactByChannelId(c echo.Context) error{
			channel_id:=  c.FormValue("channel_id")
		
			
		
				result, err := models.FetchContactByChannelId(channel_id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}