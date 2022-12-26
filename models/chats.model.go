package models

import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	"github.com/go-playground/validator"
)


type Chat struct{
	Id		int   `json:"channel_id"`
    Contact_id 	string   `json:"contact_id" validate:"required"`
}
type Allname struct{
	Id int `json:"id"`
	Name string  `json:"phone_number"`
	Phone_Number string  `json:"name"`
	Profilepic string  `json:"profilepic"`
	User_id string   `json:"user_id"`


}
//add chat, delete chat, chat by user id
// SELECT channels.name FROM channels JOIN contactschannels ON contactschannels.channel_id=channels.channel_id JOIN contacts ON contactschannels.contact_id=contacts.contact_id WHERE contacts.user_id=1 UNION SELECT groups.name FROM groups JOIN contactsgroups ON contactsgroups.group_id=groups.group_id JOIN contacts ON contactsgroups.contact_id=contacts.contact_id WHERE contacts.user_id=1 UNION SELECT contacts.name FROM chats JOIN contacts ON chats.contact_id=contacts.contact_id WHERE contacts.user_id=
func FetchChats(user_id string)(Response, error){
	var obj Allname       
	var arrObj []Allname
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	// sqlStatement := "SELECT channels.name FROM channels JOIN contactschannels ON contactschannels.channel_id=channels.channel_id JOIN contacts ON contactschannels.contact_id=contacts.contact_id WHERE contact UNION SELECT groups.name FROM groups JOIN contactsgroups ON contactsgroups.group_id=groups.group_id JOIN contacts ON contactsgroups.contact_id=contacts.contact_id UNION SELECT contacts.name FROM chats JOIN contacts ON chats.contact_id=contacts.contact_id WHERE contacts.user_id="+user_id
	sqlStatement := "SELECT  chat_id, contacts.name, contacts.phone_number,contacts.profilepic,contacts.user_id FROM chats JOIN contacts ON chats.contact_id=contacts.contact_id WHERE contacts.user_id="+user_id
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
func FetchChatsByChatId(chat_id string)(Response, error){
	var obj Allname       
	var arrObj []Allname
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT chat_id, contacts.name, contacts.phone_number,contacts.profilepic,contacts.user_id FROM chats JOIN contacts ON contacts.contact_id=chats.contact_id WHERE chats.chat_id=" + chat_id
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
func FetchAllChats(user_id string)(Response, error){
	var obj Allname       
	var arrObj []Allname
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	// sqlStatement := "SELECT channels.name FROM channels JOIN contactschannels ON contactschannels.channel_id=channels.channel_id JOIN contacts ON contactschannels.contact_id=contacts.contact_id WHERE contact UNION SELECT groups.name FROM groups JOIN contactsgroups ON contactsgroups.group_id=groups.group_id JOIN contacts ON contactsgroups.contact_id=contacts.contact_id UNION SELECT contacts.name FROM chats JOIN contacts ON chats.contact_id=contacts.contact_id WHERE contacts.user_id="+user_id
	sqlStatement := "SELECT a.id,a.name, a.phone_number, a.profilepic, a.user_id from (SELECT channels.name as name, contacts.user_id as user_id, channels.channel_id as id,contacts.profilepic as profilepic, contacts.phone_number as phone_number FROM channels JOIN contactschannels ON contactschannels.channel_id=channels.channel_id JOIN contacts ON contactschannels.contact_id=contacts.contact_id UNION SELECT groups.name as name,contacts.user_id as user_id, groups.group_id as id,contacts.profilepic as profilepic,contacts.phone_number as phone_number FROM groups JOIN contactsgroups ON contactsgroups.group_id=groups.group_id JOIN contacts ON contactsgroups.contact_id=contacts.contact_id UNION SELECT contacts.name as name,contacts.user_id as user_id, chats.chat_id as id,contacts.profilepic as profilepic,contacts.phone_number as phone_number FROM chats JOIN contacts ON chats.contact_id=contacts.contact_id)a WHERE user_id="+user_id
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


	//delete data
	func DeleteChats(id string)(Response, error){
		var res Response
		con := db.CreateCon()
		sqlStatement := "DELETE FROM chats WHERE chat_id=?"
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

//insert data
func StoreChats(contact_id string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := Chat{
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
	sqlStatement := "INSERT INTO chats(contact_id) VALUES (?)"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data=map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}
	
	result, err := stmt.Exec(contact_id)
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