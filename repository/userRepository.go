package repository


import (
	"net/http"      // Manage URL
	. "crud-rest-api/driver"
	. "crud-rest-api/domain"
	_ "github.com/go-sql-driver/mysql" // MySQL Database driver
	// "encoding/json"
)

func GetUser(w http.ResponseWriter, r *http.Request) []User {
	// Open database connection
	db := DBConn()
	defer db.Close()
	// Prepare a SQL query to select all data from database and threat errors
	selDB, err := db.Query("SELECT id_user,nama,email FROM tb_user ")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := User{}

	// Create a slice to store all data from struct
	res := []User{}

	// Read all rows from database
	for selDB.Next() {
		// Must create this variables to store temporary query
		var id int
		var name, email string

		// Scan each row storing values from the variables above and check for errors
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.IdUser = id
		n.Nama = name
		n.Email = email

		// Join each row on struct inside the Slice
		res = append(res, n)


	}
	return res
	// userResponse , errr := json.Marshal(res)
	// if errr != nil{
	// 	panic(errr.Error())
	// }
	// w.Write(userResponse)
	
}