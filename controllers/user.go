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
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
	for i, v := range slice { // Same as the previous
		println(i, v)
	}
	wellknownport := map[string]int{"http": 80, "https": 443}
	for k, v := range wellknownport {
		println(k, v)
	}
}

// Constructor, by convention new is given as staring
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
