package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// uc is like this pointer
	i := 0
	for i < 5 {
		w.Write([]byte("Anupa Suresh is a oola"))
		i++
	}

}

// Constructor, by convention new is given as staring
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
