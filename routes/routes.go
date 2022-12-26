package routes
// go mod init vp_week11_echo 
// GO111MODULE=on go get github.com/labstack/echo/v4

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"vp_week11_echo/controllers"
	"vp_week11_echo/middleware"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	name := c.Param("name")
  return c.String(http.StatusOK, "Hello, " + name)
}


func Init() *echo.Echo {
	e := echo.New()
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	
	//untuk user/profile
	e.POST("/login", controllers.CheckLogin)
	e.POST("/cekuserexist", controllers.CheckUserExist)
	e.POST("/cekcontactexist", controllers.CheckContactExist)
	e.POST("/loginsuccess", controllers.FetchAllUsers,middleware.IsAuthenticated)//byuserid, token
	e.POST("/fetchusername", controllers.FetchUsersByUsername)//byusername
	e.POST("/fetchchannelbyuserid", controllers.FetchAllChannels)//by user_id
	e.POST("/fetchchannel", controllers.FetchChannelsByChannelId)//bychannelid
	e.POST("/fetchgroupbyuserid", controllers.FetchAllGroups)//by user_id
	e.POST("/fetchgroup", controllers.FetchGroupsByGroupId)//bygroupid
	e.POST("/fetchcontactbyuserid", controllers.FetchAllContacts)//by user_id
	e.POST("/fetchcontact", controllers.FetchContactsByContactId)//byCONTACTid
	e.POST("/fetchallchatbyuserid", controllers.FetchAllChats)//by user_id
	e.POST("/fetchchatbyuserid", controllers.FetchChats)//by user_id
	e.POST("/fetchchat", controllers.FetchChatsByChatId)//by chat id
	e.POST("/fetchmessage", controllers.FetchMessagesByContactId)//by message id
	e.POST("/fetchmessagechannel", controllers.FetchMessagesByChannelId)//by message id
	e.POST("/fetchmessagegroup", controllers.FetchMessagesByGroupId)//by message id
	e.POST("/fetchcontactgroup", controllers.FetchContactByGroupId)//by message id
	e.POST("/fetchcontactchannel", controllers.FetchContactByChannelId)//by message id
	//registration, add
	e.POST("/users", controllers.StoreUsers)
	e.POST("/channels", controllers.StoreChannels)
	e.POST("/groups", controllers.StoreGroups)
	e.POST("/contacts", controllers.StoreContacts)
	e.POST("/chats", controllers.StoreChats)
	e.POST("/messages", controllers.StoreMessages)
	e.POST("/messageschannel", controllers.StoreMessagesChannels)
	e.POST("/messagesgroup", controllers.StoreMessagesGroups)
	e.POST("/contactsgroups", controllers.StoreContactsGroups)
	e.POST("/contactschannels", controllers.StoreContactsChannels)
	//edit 
	e.PATCH("/users", controllers.UpdateUsers)
	e.PATCH("/channels", controllers.UpdateChannels)
	e.PATCH("/groups", controllers.UpdateGroups)
	e.PATCH("/contacts", controllers.UpdateContacts)
	//delete 
	e.DELETE("/channels", controllers.DeleteChannels)
	e.DELETE("/groups", controllers.DeleteGroups)
	e.DELETE("/users", controllers.DeleteUsers)
	e.DELETE("/contacts", controllers.DeleteContacts)
	e.DELETE("/chats", controllers.DeleteChats)
	//bisa pengecekan dari app
	e.POST("/test-validation", controllers.TestStructValidation)
	e.POST("/test-validation-var", controllers.TestVarValidation)
	//
	return e

}

