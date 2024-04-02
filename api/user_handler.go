package api

import (
	"fmt"
	"net/http"

	"github.com/milosz94/web-go/types"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {

	var user types.User
	_ = user

	fmt.Fprintf(w, "user")
}
