package models

import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	// "github.com/go-playground/validator"
)


type Chat struct{
	Id		int   `json:"channel_id"`
    Contact_id 	int   `json:"contact_id" validate:"required"`
}
type Allname struct{
	Name string  `json:"name"`
}
//add chat, delete chat, chat by user id
// SELECT channels.name FROM channels JOIN contactschannels ON contactschannels.channel_id=channels.channel_id JOIN contacts ON contactschannels.contact_id=contacts.contact_id WHERE contacts.user_id=1 UNION SELECT groups.name FROM groups JOIN contactsgroups ON contactsgroups.group_id=groups.group_id JOIN contacts ON contactsgroups.contact_id=contacts.contact_id WHERE contacts.user_id=1 UNION SELECT contacts.name FROM chats JOIN contacts ON chats.contact_id=contacts.contact_id WHERE contacts.user_id=

func FetchAllChats(user_id string)(Response, error){
	var obj Allname       
	var arrObj []Allname
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT channels.name FROM channels JOIN contactschannels ON contactschannels.channel_id=channels.channel_id JOIN contacts ON contactschannels.contact_id=contacts.contact_id WHERE contacts.user_id=1 UNION SELECT groups.name FROM groups JOIN contactsgroups ON contactsgroups.group_id=groups.group_id JOIN contacts ON contactsgroups.contact_id=contacts.contact_id WHERE contacts.user_id=1 UNION SELECT contacts.name FROM chats JOIN contacts ON chats.contact_id=contacts.contact_id WHERE contacts.user_id="+user_id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Name)
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