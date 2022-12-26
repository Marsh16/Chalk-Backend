package controllers

import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)

//add, fetch 
func StoreMessagesChannels(c echo.Context) error{
	messages:= c.FormValue("messages")
	channel_id:= c.FormValue("channel_id")


	
		result, err := models.StoreMessagesChannels(messages,channel_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}

		func FetchMessagesByChannelId(c echo.Context) error{
			channel_id:=  c.FormValue("channel_id")
		
			
		
				result, err := models.FetchMessagesByChannelId(channel_id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}