package main

import (
	"net/http"

	"github.com/jkss2110/CMSservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
