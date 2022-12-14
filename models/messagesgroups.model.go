package models

// import (
// 	"vp_week11_echo/db"
// 	"net/http"
// 	//"encoding/json"
// 	"github.com/go-playground/validator"
// )


type MessageGroup struct{
	Id		int   `json:"messagesgroup_id"`
    Messages 	string   `json:"messages" validate:"required"`
	Group_id int   `json:"group_id" validate:"required"`
}