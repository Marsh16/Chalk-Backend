package models

import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	// "database/sql"
	
	// "fmt"
	"github.com/go-playground/validator"
)
func FetchUsersByUsername(username string)(Response, error){//error
	var obj string
	// var arrObj []User
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT user_id FROM users WHERE username = ?"
	err := con.QueryRow(sqlStatement, username).Scan(
		&obj,
	)
	
	

	if err != nil{
		return res,err
	}

	res.Status = http.StatusOK
	res.Message="Success"
	res.Data = obj

	return res, nil
	

}
//Read all
func FetchAllUsers(user_id string)(Response, error){
	var obj User
	var arrObj []User
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM users where user_id="+ user_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Name, &obj.Username,&obj.Email,&obj.Phone_Number, &obj.Dateofbirth,&obj.Profilepic, &obj.Password)
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
func StoreUsers(name string, username string, email string, phone_number string,dateofbirth string,profilepic string,password string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := User{
		Name: name,
		Username: username,
		Email: email,
		Phone_Number: phone_number,
		Dateofbirth: dateofbirth,
		Profilepic: profilepic,
		Password: password,
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
	sqlStatement := "INSERT INTO users(name,username,email,phone_number,dateofbirth,profilepic,password) VALUES (?,?,?,?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(name,username,email,phone_number,dateofbirth,profilepic,password)
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
//update data
func UpdateUsers(id string,name string, username string, email string, phone_number string,dateofbirth string,profilepic string,password string)(Response, error){
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE users SET name=?,username=?,email=?,phone_number=?,dateofbirth=?,profilepic=?,password=? WHERE user_id=?"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		return res, err
	}
	
	result, err := stmt.Exec(name,username,email,phone_number,dateofbirth,profilepic,password,id)
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

	//delete data
	func DeleteUsers(id string)(Response, error){
		var res Response
		con := db.CreateCon()
		sqlStatement := "DELETE FROM users WHERE user_id=?"
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