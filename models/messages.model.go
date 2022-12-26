package models

// import (
// 	"vp_week11_echo/db"
// 	"net/http"
// 	//"encoding/json"
// 	"github.com/go-playground/validator"
// )


type Message struct{
	Id		int   `json:"messages_id"`
    Messages 	string   `json:"messages" validate:"required"`
    Contact_id int   `json:"contact_id" validate:"required"`
}

//add, fecth by user id , fetch messages by contact id, 

func FetchAllContact(user_id string)(Response, error){
	var obj Contact
	var arrObj []Contact
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT * from contacts where user_id= "+user_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(	&obj.Id, &obj.Name, &obj.Phone_Number,&obj.Profilepic,&obj.User_id )
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
func StoreContacts(name string, phone_number string, profilepic string, user_id string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := Contact{
		Name: name,
		Phone_Number: phone_number,
		Profilepic: profilepic,
		User_id:user_id,
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
	sqlStatement := "INSERT INTO contacts(name,phone_number,profilepic,user_id) VALUES (?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(name,phone_number,profilepic,user_id)
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
func FetchContactsByContactId(contact_id string)(Response, error){
	var obj Contact
	var arrObj []Contact
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from contacts WHERE contact_id=" + contact_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Phone_Number,&obj.Profilepic, &obj.User_id)
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




	


