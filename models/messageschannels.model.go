package models

import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	"github.com/go-playground/validator"
)


type MessageChannel struct{
	Id		int   `json:"messageschannels_id"`
    Messages 	string   `json:"messages" validate:"required"`
	Channel_id string   `json:"channel_id" validate:"required"`
}


//add, fecth by user id , fetch messages by contact id, 

func FetchMessagesByChannelId(channel_id string)(Response, error){
	var obj MessageChannel
	var arrObj []MessageChannel
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT * from messageschannels where channel_id= "+channel_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Messages, &obj.Channel_id)
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
func StoreMessagesChannels(messages string, channel_id string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := MessageChannel{
		Messages: messages,
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
	sqlStatement := "INSERT INTO messageschannels(messages,channel_id) VALUES (?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(messages,channel_id)
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




	


