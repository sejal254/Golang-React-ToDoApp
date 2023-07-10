package main

import (
	"fmt"
	//"golang-react-todo/router"
	"log"
	"net/http"

	"github.com/sejal254/Go-React-ToDoApp/router"
)

func main() {

	r := router.Router()
	fmt.Println("Staring the Server on port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
