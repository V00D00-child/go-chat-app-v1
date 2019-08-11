package handlers

import (
	"fmt"
	"net/http"
	"realtime-chat-go-api/core"
	"realtime-chat-go-api/models"
	"realtime-chat-go-api/services"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// TODO: add web and event logs to this file and handle the check token function.

// TokenHandler ... When a user makes a request to get a token they must pass their username
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	// Create access logger
	currentDate := time.Now()
	writer := models.LoggingResponseWriter{}
	accessLog := models.NewWebLogEntry(r, &writer, currentDate)

	// Obtain username from URL
	vars := mux.Vars(r)
	key := vars["username"]

	// Check to make sure user is in the database
	tempUser, err := services.GetUserByUserName(strings.ToLower(key))
	if err != nil {
		accessLog.Status = 500
		models.WriteWebLog(accessLog, "accessLogs.json")
		models.FailureHandlerMessage("Failed to get user", w, accessLog.Status, 6)
		return
	}

	// Extract password  from users
	user := services.DbUserToUser(tempUser)

	// Create a JWT token for the user
	tokenString, _ := core.CreateToken(user)

	// return the token to the user
	w.Header().Set("Authorization", "Bearer:"+tokenString)
}

/*RefreshTokenHandler ...
/token GET |POST
*/
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Create access logger
	// currentDate := time.Now()
	// writer := models.LoggingResponseWriter{}
	// accessLog := models.NewWebLogEntry(r, &writer, currentDate)

	// Get user token
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		models.FailureHandlerMessage("Missing Authorization header", w, http.StatusUnauthorized, 6)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer:", "", 1)

	// Check to token
	claim, err := core.CheckToken(tokenString)
	if err != nil {
		models.FailureHandlerMessage("Token in not valid", w, http.StatusUnauthorized, 6)
	}

	if r.Method == http.MethodGet {

		// Check to see if token needs to be renewed
		newToken, isRefresh := core.RenewToken(claim)
		if isRefresh == true {
			// Return the new token in the response body
			w.Header().Set("Authorization", "Bearer:"+newToken)
			models.SuccessHandlerMessage("Token has been renewed", "", w, 200, 0)
		} else {
			models.FailureHandlerMessage("Token does not need to be renewed yet", w, http.StatusBadRequest, 0)
			return
		}
	} else if r.Method == http.MethodPost {
		if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) <= 0 {
			fmt.Println(time.Unix(claim.ExpiresAt, 0).Sub(time.Now()))
			models.FailureHandlerMessage("Token in not valid", w, http.StatusUnauthorized, 6)
		}
	}

}
