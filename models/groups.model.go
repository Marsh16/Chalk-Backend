package models

import (
	"vp_week11_echo/db"
	"net/http"
// 	//"encoding/json"
	"github.com/go-playground/validator"
 )


type Group struct{
	Id		int   `json:"group_id"`
    Name 	string   `json:"name" validate:"required"`
    Description 	string   `json:"description" validate:"required"`
	Profilepic	string   `json:"profilepic"`
	// Contact_id int   `json:"contact_id" validate:"required"`
}

// SELECT groups.name, channels.name, contacts.name FROM groups, channels, chats JOIN contacts ON chats.contact_id=contacts.contact_id WHERE contacts.user_id=1;
func FetchAllGroups(user_id string)(Response, error){
	var obj Group
	var arrObj []Group
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT groups.group_id,groups.name,groups.description,groups.profilepic  from groups JOIN contactsgroups ON groups.group_id=contactsgroups.group_id JOIN contacts ON contacts.contact_id=contactsgroups.contact_id WHERE contacts.user_id= "+user_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Name, &obj.Description,&obj.Profilepic, )
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
func StoreGroups(name string, description string, profilepic string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := Group{
		Name: name,
		Description: description,
		Profilepic: profilepic,
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
	sqlStatement := "INSERT INTO groups(name,description,profilepic) VALUES (?,?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(name,description,profilepic)
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
func FetchGroupsByGroupId(group_id string)(Response, error){
	var obj Group
	var arrObj []Group
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from groups WHERE group_id=" + group_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Description,&obj.Profilepic)
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
	//delete data
	func DeleteGroups(id string)(Response, error){
		var res Response
		con := db.CreateCon()
		sqlStatement := "DELETE FROM groups WHERE group_id=?"
		stmt, err := con.Prepare(sqlStatement)
		
		if err!= nil{
			return res, err
		}
		
		result, err := stmt.Exec(id)
		if err!= nil{
			return res, err
		}
		
		rowsAffected, err := result.RowsAffected()
		
		if err!= nil{
			return res, err
		}
		res.Status = http.StatusOK
			res.Message="Success"
			res.Data = map[string]int64{
		"rows_affected":rowsAffected,
		
			}
			return res,nil
		
		}



		func UpdateGroups(id string,name string, description string, profilepic string)(Response, error){
			var res Response
			con := db.CreateCon()
			sqlStatement := "UPDATE groups SET name=?,description=?,profilepic=? WHERE group_id=?"
			stmt, err := con.Prepare(sqlStatement)
			
			if err!= nil{
				return res, err
			}
			
			result, err := stmt.Exec(name,description,profilepic,id)
			if err!= nil{
				return res, err
			}
			
			rowsAffected, err := result.RowsAffected()
			
			if err!= nil{
				return res, err
			}
			res.Status = http.StatusOK
				res.Message="Success"
				res.Data = map[string]int64{
			"rows_affected":rowsAffected,
			
				}
				return res,nil
			
			}
