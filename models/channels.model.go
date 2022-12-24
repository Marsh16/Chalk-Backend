package models


import (
	"vp_week11_echo/db"
	"net/http"
	"github.com/go-playground/validator"
	//"encoding/json"
)


type Channel struct{
	Id		int   `json:"channel_id"`
    Name 	string   `json:"name" validate:"required"`
    Description 	string   `json:"description" validate:"required"`
	Profilepic	string   `json:"profilepic"`
	//ContactChannel_id int   `json:"contactchannel_id" validate:"required"`

}

func FetchAllChannels(user_id string)(Response, error){
	var obj Channel
	var arrObj []Channel
	var res Response
	con:= db.CreateCon()
	// , contacts.user_id 

	sqlStatement := "SELECT channels.channel_id,channels.name,channels.description,channels.profilepic from channels JOIN contactschannels ON channels.channel_id=contactschannels.channel_id JOIN contacts ON contacts.contact_id=contactschannels.contact_id WHERE contacts.user_id= "+user_id
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
func StoreChannels(name string, description string, profilepic string)(Response, error){
	var res Response
	
	v := validator.New()
	
	usr := Channel{
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
	sqlStatement := "INSERT INTO channels(name,description,profilepic) VALUES (?,?,?)"
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
func FetchChannelsByChannelId(channel_id string)(Response, error){
	var obj Channel
	var arrObj []Channel
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from channels WHERE channel_id=" + channel_id
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
	func DeleteChannels(id string)(Response, error){
		var res Response
		con := db.CreateCon()
		sqlStatement := "DELETE FROM channels WHERE channel_id=?"
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



		func UpdateChannels(id string,name string, description string, profilepic string)(Response, error){
			var res Response
			con := db.CreateCon()
			sqlStatement := "UPDATE channels SET name=?,description=?,profilepic=? WHERE channel_id=?"
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
