package models

// import (
// 	"vp_week11_echo/db"
// 	"net/http"
// 	//"encoding/json"
// 	"github.com/go-playground/validator"
// )


type Chat struct{
	Id		int   `json:"channel_id"`
    Contact_id 	int   `json:"contact_id" validate:"required"`
}
