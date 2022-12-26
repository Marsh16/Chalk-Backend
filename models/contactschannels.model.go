package models

import (
	"vp_week11_echo/db"
	"net/http"
	// "fmt"
	//"encoding/json"
	// "database/sql"
	"github.com/go-playground/validator"
)


type ContactChannel struct{
	Id		int   `json:"contactchannel_id"`
    Contact_id 	string   `json:"contact_id" validate:"required"`
    Channel_id 	string   `json:"channel_id" validate:"required"`
}


//fetch where channelid, add 
func FetchContactByChannelId(channel_id string)(Response, error){
	var obj ContactChannel
	var arrObj []ContactChannel
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT * from contactschannels where channel_id= "+channel_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Contact_id, &obj.Channel_id)
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
//insert data
func StoreContactsChannels(contact_id string, channel_id string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := ContactChannel{
		Contact_id: contact_id,
		Channel_id: channel_id,
	}
	
	err := v.Struct(usr)
	if err!= nil{
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	con := db.CreateCon()
	sqlStatement := "INSERT INTO contactschannels(contact_id,channel_id) VALUES (?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(contact_id,channel_id)
	if err!= nil{
		return res, err
	}
	
	lastInsertedID, err := result.LastInsertId()
	
	if err!= nil{
		return res, err
	}
	res.Status = http.StatusOK
		res.Message="Success"
		res.Data = map[string]int64{
	"last_inserted_id":lastInsertedID,
	
		}
		return res,nil
	
	}
