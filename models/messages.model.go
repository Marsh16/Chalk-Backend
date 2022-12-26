package models

import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	"github.com/go-playground/validator"
)


type Message struct{
	Id		int   `json:"messages_id"`
    Messages 	string   `json:"messages" validate:"required"`
    Contact_id string   `json:"contact_id" validate:"required"`
}

//add, fecth by user id , fetch messages by contact id, 

func FetchMessagesByContactId(contact_id string)(Response, error){
	var obj Message
	var arrObj []Message
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT * from messages where contact_id= "+contact_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Messages, &obj.Contact_id)
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
func StoreMessages(messages string, contact_id string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := Message{
		Messages: messages,
		Contact_id: contact_id,
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
	sqlStatement := "INSERT INTO messages(messages,contact_id) VALUES (?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(messages,contact_id)
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




	


