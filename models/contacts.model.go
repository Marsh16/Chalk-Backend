package models

// import (
// 	"vp_week11_echo/db"
// 	"net/http"
// 	//"encoding/json"
// 	"github.com/go-playground/validator"
// )


type Contact struct{
	Id		int   `json:"contact_id"`
    Name 	string   `json:"name" validate:"required"`
    Phone_Number 	string   `json:"phone_number" validate:"required"`
	Profilepic	string   `json:"profilepic"`
	User_id int   `json:"user_id" validate:"required"`
}
