package main

import (
	"net/http"

	"github.com/shishanksingh2015/GoWebService/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
