package main

import "fmt"
import "net/http"
// import . "crud-rest-api/driver"
// import . "crud-rest-api/domain"
import . "crud-rest-api/service"


func main(){


	fmt.Print("listen at 9000")
	http.HandleFunc("/",GetUserService)

	// Start the server on port 9000
	http.ListenAndServe(":9000", nil)
}