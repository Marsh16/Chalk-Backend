package controllers


import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)
func UpdateChannels(c echo.Context) error{
	id:=  c.FormValue("channel_id")
	name:= c.FormValue("name")
description:= c.FormValue("description")
profilepic:= c.FormValue("profilepic")

		result, err := models.UpdateChannels(id,name,description,profilepic)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func StoreChannels(c echo.Context) error{
	name:= c.FormValue("name")
	description:= c.FormValue("description")
	profilepic:= c.FormValue("profilepic")

	
		result, err := models.StoreChannels(name,description,profilepic)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func FetchAllChannels(c echo.Context) error{
	user_id:=  c.FormValue("user_id")
	

		result, err := models.FetchAllChannels(user_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func DeleteChannels(c echo.Context) error{
	id:= c.FormValue("channel_id")
			
	result, err := models.DeleteChannels(id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}

func FetchChannelsByChannelId(c echo.Context) error{
	channel_id:=  c.FormValue("channel_id")

	

		result, err := models.FetchChannelsByChannelId(channel_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}