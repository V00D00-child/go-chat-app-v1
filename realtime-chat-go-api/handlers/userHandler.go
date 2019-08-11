package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"realtime-chat-go-api/models"
	"realtime-chat-go-api/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/*UsersHandler ...
/users GET |POST
*/
func UsersHandler(w http.ResponseWriter, r *http.Request) {

	// Create access logger
	currentDate := time.Now()
	writer := models.LoggingResponseWriter{}
	accessLog := models.NewWebLogEntry(r, &writer, currentDate)

	// Setup Response
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Handlers GET |POST
	if r.Method == http.MethodGet {

		dbUser, err := services.GetUsers()
		if err != nil {
			accessLog.Status = 500
			models.WriteWebLog(accessLog, "accessLogs.json")
			models.FailureHandlerMessage("Failed to get all users", w, accessLog.Status, 6)
			return
		}

		// Extract password  from users
		users := services.DbUsersToUsers(dbUser)

		// Write response
		accessLog.Status = 200
		models.WriteWebLog(accessLog, "accessLogs.json")
		models.SuccessHandlerMessage("Getting all users", users, w, accessLog.Status, 0)

	} else if r.Method == http.MethodPost {

		// Get the request body user
		var tempUser models.DbUser
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &tempUser)

		//  Add user to data base
		err := services.CreateUser(tempUser)
		if err != nil {
			accessLog.Status = 500
			models.WriteWebLog(accessLog, "accessLogs.json")
			models.FailureHandlerMessage("Failed to create a user", w, accessLog.Status, 6)
			return
		}

		// Write response
		accessLog.Status = 201
		models.WriteWebLog(accessLog, "accessLogs.json")
		models.SuccessHandlerMessage("Created new user", nil, w, accessLog.Status, 0)
	}

}

/*UserHandler ...
/users/{id} GET|DELETE |
*/
func UserHandler(w http.ResponseWriter, r *http.Request) {

	// Create access logger
	currentDate := time.Now()
	writer := models.LoggingResponseWriter{}
	accessLog := models.NewWebLogEntry(r, &writer, currentDate)

	// SetupResponse
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodGet {

		// Obtain id from URL
		vars := mux.Vars(r)
		key := vars["id"]

		keyInt, err := strconv.Atoi(key)
		if err != nil {
			accessLog.Status = 500
			models.WriteWebLog(accessLog, "accessLogs.json")
			models.FailureHandlerMessage("Internal server error", w, accessLog.Status, 6)
			return
		}

		tempUser, err := services.GetUserByID(keyInt)
		if err != nil {
			accessLog.Status = 500
			fmt.Println(err)
			models.WriteWebLog(accessLog, "accessLogs.json")
			models.FailureHandlerMessage("Failed to get user", w, accessLog.Status, 6)
			return
		}

		// Extract password  from users
		user := services.DbUserToUser(tempUser)

		// Write response
		accessLog.Status = 200
		models.WriteWebLog(accessLog, "accessLogs.json")
		models.SuccessHandlerMessage("Getting single user", user, w, accessLog.Status, 0)

	} else if r.Method == http.MethodDelete {

		// Obtain id from URL
		vars := mux.Vars(r)
		key := vars["id"]

		err := services.DeleteUserByID(key)
		if err != nil {
			accessLog.Status = 500
			models.WriteWebLog(accessLog, "accessLogs.json")
			models.FailureHandlerMessage("Failed to delete user", w, accessLog.Status, 6)
			return
		}

		// Write response
		accessLog.Status = 204
		models.WriteWebLog(accessLog, "accessLogs.json")
		models.SuccessHandlerMessage("Deleting a user", nil, w, accessLog.Status, 0)
	}
}
