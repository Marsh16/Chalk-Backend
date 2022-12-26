package models

import (
	"vp_week11_echo/db"
	"net/http"
	// "fmt"
	//"encoding/json"
	// "database/sql"
	"github.com/go-playground/validator"
)


type ContactGroup struct{
	Id		int   `json:"contactgroup_id"`
    Contact_id 	string   `json:"contact_id" validate:"required"`
    Group_id 	string   `json:"group_id" validate:"required"`
}

func FetchContactByGroupId(group_id string)(Response, error){
	var obj ContactGroup
	var arrObj []ContactGroup
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT * from contactsgroups where group_id= "+group_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Contact_id, &obj.Group_id)
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
func StoreContactsGroups(contact_id string, group_id string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := ContactGroup{
		Contact_id: contact_id,
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
	sqlStatement := "INSERT INTO contactsgroups(contact_id,group_id) VALUES (?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(contact_id,group_id)
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
