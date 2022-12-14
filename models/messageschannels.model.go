package models

// import (
// 	"vp_week11_echo/db"
// 	"net/http"
// 	//"encoding/json"
// 	"github.com/go-playground/validator"
// )


type MessageChannel struct{
	Id		int   `json:"messageschannels_id"`
    Messages 	string   `json:"messages" validate:"required"`
	Channel_id int   `json:"channel_id" validate:"required"`
}
