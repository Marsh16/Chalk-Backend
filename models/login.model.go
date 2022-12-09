package models

import (
	"vp_week11_echo/db"
	//"encoding/json"
	"database/sql"
	"vp_week11_echo/helpers"
	"fmt"
	"net/http"
)

type User struct{
	// Id			int `json:"id"`
	// Username	string `json:"username"`
	// Password	string `json:"password"`
	// Email		string `json:"email"`
	Id		int   `json:"id"`
    Name 	string   `json:"name" validate:"required"`
    Username 	string   `json:"username" validate:"required"`
    Email	string   `json:"email" validate:"required"`
	Phone_Number	string   `json:"phone_number" validate:"required"`
	Dateofbirth	string   `json:"dateofbirth" validate:"required"`
	Password string   `json:"password" validate:"required"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"
	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Name, &obj.Username,&obj.Email,&obj.Phone_Number, &obj.Dateofbirth, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Print("Username not found!")
		return false, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return false, err
	}

	return true, nil
}

//Read all
func FetchAllUsers()(Response, error){
	var obj User
	var arrObj []User
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM users"
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Name, &obj.Username,&obj.Email,&obj.Phone_Number, &obj.Dateofbirth, &obj.Password)
		if err != nil{
			return res,err
		}
		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message="Success"
	res.Data = arrObj
	//b, err := json.MarshalIndent(res, "", "    ")

	return res, nil
}