package service 

import (
	"net/http"      // Manage URL
	// . "crud-rest-api/driver"
	. "crud-rest-api/repository"
	_ "github.com/go-sql-driver/mysql" // MySQL Database driver
	"encoding/json"
)

func GetUserService(w http.ResponseWriter, r *http.Request){
	userResponseFromRepo := GetUser(w,r)
	userResponse , errr := json.Marshal(userResponseFromRepo)
	if errr != nil{
		panic(errr.Error())
	}
	w.Write(userResponse)
}