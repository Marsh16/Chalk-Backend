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
}

func FetchAllChannels()(Response, error){
	var obj Channel
	var arrObj []Channel
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM channels"
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Name, &obj.Description,&obj.Profilepic,&obj.Contact_id)
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
// func StoreUsers(name string, description string, profilepic string, contact_id string)(Response, error){
// 	var res Response
	
// 	v := validator.New()
	
// 	usr := Channel{
// 		Name: name,
// 		Description: description,
// 		Profilepic: profilepic,
// 		Contact_id: contact_id,
// 	}
	
// 	err := v.Struct(usr)
// 	if err!= nil{
// 		res.Status = http.StatusBadRequest
// 		res.Message = "Error"
// 		res.Data=map[string]string{
// 			"errors": err.Error(),
// 		}
// 		return res, err
// 	}
	
// 	con := db.CreateCon()
// 	sqlStatement := "INSERT INTO users(name,username,email,phone_number,dateofbirth,profilepic,password) VALUES (?,?,?,?,?,?,?)"
// 	stmt, err := con.Prepare(sqlStatement)
	
// 	if err!= nil{
// 		res.Status = http.StatusInternalServerError
// 		res.Message = "Error"
// 		res.Data=map[string]string{
// 			"errors": err.Error(),
// 		}
// 		return res, err
// 	}
	
// 	result, err := stmt.Exec(name,username,email,phone_number,dateofbirth,profilepic,password)
// 	if err!= nil{
// 		return res, err
// 	}
	
// 	lastInsertedID, err := result.LastInsertId()
	
// 	if err!= nil{
// 		return res, err
// 	}
// 	res.Status = http.StatusOK
// 		res.Message="Success"
// 		res.Data = map[string]int64{
// 	"last_inserted_id":lastInsertedID,
	
// 		}
// 		return res,nil
	
// 	}