package controllers


import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)
func UpdateGroups(c echo.Context) error{
	id:=  c.FormValue("group_id")
	name:= c.FormValue("name")
description:= c.FormValue("description")
profilepic:= c.FormValue("profilepic")

		result, err := models.UpdateGroups(id,name,description,profilepic)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func StoreGroups(c echo.Context) error{
	name:= c.FormValue("name")
	description:= c.FormValue("description")
	profilepic:= c.FormValue("profilepic")

	
		result, err := models.StoreGroups(name,description,profilepic)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func FetchAllGroups(c echo.Context) error{
	user_id:=  c.FormValue("user_id")
	

		result, err := models.FetchAllGroups(user_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func DeleteGroups(c echo.Context) error{
	id:= c.FormValue("group_id")
			
	result, err := models.DeleteChannels(id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}

func FetchGroupsByGroupId(c echo.Context) error{
	group_id:=  c.FormValue("group_id")

	

		result, err := models.FetchGroupsByGroupId(group_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}