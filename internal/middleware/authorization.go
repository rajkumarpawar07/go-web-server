package middleware

import (
	"errors"
	"go_tutorials/webserver/cmd/api"
	"net/http"
	log "github.com/sirupsen/logrus"

	"go_tutorials/webserver/internal/tools"
)


var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the username from the request
		username := r.URL.Query().Get("username")
		// get the token from the request
		token := r.Header.Get("Authorization")
		var err error
		// check if the username is empty
		if username == "" || token == ""{
			// if the username is empty, return an error
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		// check if the username and token are valid
		var loginDetails *tools.LoginDetails
		loginDetails= (*database).GetUserLoginDetails(username)
		if (loginDetails == nil || (token!=(*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// if the username and token are valid, call the next handler
		next.ServeHTTP(w, r)
	})
}