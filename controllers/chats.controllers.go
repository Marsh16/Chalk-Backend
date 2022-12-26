package controllers
import (
	"github.com/labstack/echo/v4"
"vp_week11_echo/models"
"net/http"
)
func FetchAllChats(c echo.Context) error{
	user_id:=  c.FormValue("user_id")
	

		result, err := models.FetchAllChats(user_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
		func FetchChats(c echo.Context) error{
			user_id:=  c.FormValue("user_id")
			
		
				result, err := models.FetchChats(user_id)
				
				if err != nil{
					return c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				}
				return c.JSON(http.StatusOK, result)
				}
func DeleteChats(c echo.Context) error{
		id:= c.FormValue("chat_id")
					
			result, err := models.DeleteChannels(id)
						
						if err != nil{
							return c.JSON(http.StatusInternalServerError,
							map[string]string{"message": err.Error()})
						}
						return c.JSON(http.StatusOK, result)
						}
func FetchChatsByChatId(c echo.Context) error{
	chat_id:=  c.FormValue("chat_id")

	

		result, err := models.FetchChatsByChatId(chat_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}
func StoreChats(c echo.Context) error{
	contact_id:= c.FormValue("contact_id")
	
	
		result, err := models.StoreChats(contact_id)
		
		if err != nil{
			return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
		}