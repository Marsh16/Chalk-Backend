package models

import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	"github.com/go-playground/validator"
)


type MessageGroup struct{
	Id		int   `json:"messagesgroup_id"`
    Messages 	string   `json:"messages" validate:"required"`
	Group_id string   `json:"group_id" validate:"required"`
}

//add, fecth by user id , fetch messages by contact id, 

func FetchMessagesByGroupId(group_id string)(Response, error){
	var obj MessageGroup
	var arrObj []MessageGroup
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT * from messagesgroups where group_id= "+group_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Messages, &obj.Group_id)
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
func StoreMessagesGroups(messages string, group_id string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := MessageGroup{
		Messages: messages,
		Group_id: group_id,
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
	sqlStatement := "INSERT INTO messagesgroups(messages,group_id) VALUES (?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(messages,group_id)
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




	
