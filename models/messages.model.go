package models

// import (
// 	"vp_week11_echo/db"
// 	"net/http"
// 	//"encoding/json"
// 	"github.com/go-playground/validator"
// )


type Message struct{
	Id		int   `json:"messages_id"`
    Messages 	string   `json:"messages" validate:"required"`
    Contact_id int   `json:"contact_id" validate:"required"`
}
