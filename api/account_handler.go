package api

import (
	"fmt"
	"net/http"

	"github.com/milosz94/web-go/types"
)

func HandleGetAccount(w http.ResponseWriter, r *http.Request) {

	var account types.Account
	_ = account

	fmt.Fprintf(w, "account")
}
