package models

import (
	"vp_week11_echo/db"
	"net/http"
	//"encoding/json"
	"github.com/go-playground/validator"
)
	


type Mahasiswa struct {
    Id		int   `json:"id"`
    Nim 	string   `json:"nim" validate:"required,numeric,len=13"`
    Name 	string   `json:"name" validate:"required"`
    Gender	string   `json:"gender" validate:"required"`
	Fakultas	string   `json:"fakultas" validate:"required"`
	Prodi	string   `json:"prodi" validate:"required"`
}

//Read all
func FetchAllMahasiswa()(Response, error){
	var obj Mahasiswa
	var arrObj []Mahasiswa
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM mahasiswa"
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(&obj.Id,&obj.Nim,&obj.Name,&obj.Gender,&obj.Fakultas,&obj.Prodi)
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
//insert data
func StoreMahasiswa(nim string, name string, gender string, fakultas string, prodi string)(Response, error){
var res Response

v := validator.New()

mhs := Mahasiswa{
	Nim: nim,
	Name: name,
	Gender: gender,
	Fakultas: fakultas,
	Prodi: prodi,
}

err := v.Struct(mhs)
if err!= nil{
	res.Status = http.StatusBadRequest
	res.Message = "Error"
	res.Data=map[string]string{
		"errors": err.Error(),
	}
	return res, err
}

con := db.CreateCon()
sqlStatement := "INSERT INTO mahasiswa(nim,name,gender,fakultas,prodi) VALUES (?,?,?,?,?)"
stmt, err := con.Prepare(sqlStatement)

if err!= nil{
	res.Status = http.StatusInternalServerError
	res.Message = "Error"
	res.Data=map[string]string{
		"errors": err.Error(),
	}
	return res, err
}

result, err := stmt.Exec(nim, name, gender, fakultas, prodi)
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
func UpdateMahasiswa(nim string, name string, gender string, fakultas string, prodi string)(Response, error){
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE mahasiswa SET name=?,gender=?,fakultas=?,prodi=? WHERE nim=?"
	stmt, err := con.Prepare(sqlStatement)
	
	if err!= nil{
		return res, err
	}
	
	result, err := stmt.Exec(name, gender, fakultas, prodi, nim)
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
	func DeleteMahasiswa(nim string)(Response, error){
		var res Response
		con := db.CreateCon()
		sqlStatement := "DELETE FROM mahasiswa WHERE nim=?"
		stmt, err := con.Prepare(sqlStatement)
		
		if err!= nil{
			return res, err
		}
		
		result, err := stmt.Exec(nim)
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