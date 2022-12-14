package models

// import (
// 	"vp_week11_echo/db"
// 	"net/http"
// 	//"encoding/json"
// 	"github.com/go-playground/validator"
// )


type Group struct{
	Id		int   `json:"group_id"`
    Name 	string   `json:"name" validate:"required"`
    Description 	string   `json:"description" validate:"required"`
	Profilepic	string   `json:"profilepic"`
	Contact_id int   `json:"contact_id" validate:"required"`
}
