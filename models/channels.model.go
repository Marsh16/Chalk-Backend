package models


import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	// "github.com/go-playground/validator"
)


type Channel struct{
	Id		int   `json:"channel_id"`
    Name 	string   `json:"name" validate:"required"`
    Description 	string   `json:"description" validate:"required"`
	Profilepic	string   `json:"profilepic"`
	Contact_id int   `json:"contact_id" validate:"required"`
	User_id int `json:"user_id" validate:"required"`
}

func FetchAllChannels(user_id string)(Response, error){
	var obj Channel
	var arrObj []Channel
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT channels.channel_id, channels.name, channels.description, channels.profilepic, channels.contact_id, contacts.user_id from channels JOIN contacts ON channels.contact_id=contacts.contact_id WHERE contacts.user_id= "+user_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Name, &obj.Description,&obj.Profilepic,&obj.Contact_id, &obj.User_id)
		if err != nil{
			return res,err
		}
		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message="Success"
	res.Data = arrObj

	return res, nil
}




