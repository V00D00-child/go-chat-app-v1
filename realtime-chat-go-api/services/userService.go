package services

import (
	"database/sql"
	"realtime-chat-go-api/models"
	"strconv"
	"strings"
)

// GetUsers ...
func GetUsers() ([]models.DbUser, error) {

	// connect to database a query all users
	db, err := sql.Open("mysql", "root:az0f4c3f1jqc0um8@tcp(localhost:3306)/chattie")
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to connect to database", "GetUsers()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return nil, err
	}
	defer db.Close()

	// Execute the SELECT query
	results, err := db.Query("SELECT User_ID, User_Name, User_Email,User_Password, Active,Online_Status FROM App_User")
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to execute query", "GetUsers()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return nil, err
	}

	var tempUsers []models.DbUser
	i := 0
	for results.Next() {
		var user models.DbUser
		err = results.Scan(&user.ID, &user.UserName, &user.UserEmail, &user.UserPassword, &user.Active, &user.OnlineStatus)
		if err != nil {
			ev := models.NewEventLogEntry("Error when trying to scan database results into user struct", "GetUsers()", "User service error", err)
			models.WriteEventLog(ev, "eventLogs.json")
			return nil, err
		}
		tempUsers = append(tempUsers, user)
		i++
	}
	return tempUsers, nil
}

// GetUserByID ...
func GetUserByID(id int) (models.DbUser, error) {
	var user models.DbUser

	// connect to database a query all users
	db, err := sql.Open("mysql", "root:az0f4c3f1jqc0um8@tcp(localhost:3306)/chattie")
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to connect to database", "GetUserByID()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return user, err
	}
	defer db.Close()

	// Execute the query
	err = db.QueryRow("SELECT User_ID, User_Name, User_Email,User_Password, Active,Online_Status FROM App_User where User_ID = ?", id).Scan(&user.ID, &user.UserName, &user.UserEmail, &user.UserPassword, &user.Active, &user.OnlineStatus)
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to execute query", "GetUserByID()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return user, err
	}
	return user, nil
}

// GetUserByUserName ...
func GetUserByUserName(username string) (models.DbUser, error) {
	var user models.DbUser

	// connect to database a query all users
	db, err := sql.Open("mysql", "root:az0f4c3f1jqc0um8@tcp(localhost:3306)/chattie")
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to connect to database", "GetUserByUserName()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return user, err
	}
	defer db.Close()

	// Execute the query
	err = db.QueryRow("SELECT User_ID, User_Name, User_Email,User_Password, Active,Online_Status FROM App_User where User_Name = ?", username).Scan(&user.ID, &user.UserName, &user.UserEmail, &user.UserPassword, &user.Active, &user.OnlineStatus)
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to execute query", "GetUserByUserName()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return user, err
	}
	return user, nil
}

// CreateUser ...
func CreateUser(u models.DbUser) error {

	// connect to database a query all users
	db, err := sql.Open("mysql", "root:az0f4c3f1jqc0um8@tcp(localhost:3306)/chattie")
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to connect to database", "CreateUser()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return err
	}
	defer db.Close()

	id := strconv.Itoa(u.ID)
	active := strconv.Itoa(u.Active)
	online := strconv.Itoa(u.OnlineStatus)

	query := `INSERT INTO App_User (User_ID, User_Name, User_Email,User_Password, Active,Online_Status) VALUES ( {id}, '{username}', '{email}', '{password}',{active},{online})`
	query = strings.Replace(query, "{id}", id, 1)
	query = strings.Replace(query, "{username}", u.UserName, 1)
	query = strings.Replace(query, "{email}", u.UserEmail, 1)
	query = strings.Replace(query, "{password}", u.UserPassword, 1)
	query = strings.Replace(query, "{active}", active, 1)
	query = strings.Replace(query, "{online}", online, 1)

	// Perform a INSERT Query
	insert, err := db.Query(query)
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to execute query", "CreateUser()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	return nil
}

// DeleteUserByID ...
func DeleteUserByID(id string) error {

	// connect to database a query all users
	db, err := sql.Open("mysql", "root:az0f4c3f1jqc0um8@tcp(localhost:3306)/chattie")
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to connect to database", "DeleteUserByID()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return err
	}
	defer db.Close()

	query := `DELETE FROM App_User WHERE User_ID= {id}`
	query = strings.Replace(query, "{id}", id, 1)

	// Perform a INSERT Query
	insert, err := db.Query(query)
	if err != nil {
		ev := models.NewEventLogEntry("Error when trying to execute query", "DeleteUserByID()", "User service error", err)
		models.WriteEventLog(ev, "eventLogs.json")
		return err
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	return nil
}

// DbUsersToUsers ... Returns a list of users with the password  extracted
func DbUsersToUsers(dbUsers []models.DbUser) []models.User {
	var ret = make([]models.User, len(dbUsers))
	for i, x := range dbUsers {
		ret[i].UserEmail = x.UserEmail
		ret[i].ID = x.ID
		ret[i].UserName = x.UserName
		ret[i].Active = x.Active
		ret[i].OnlineStatus = x.OnlineStatus
	}
	return ret
}

// DbUserToUser ...  Returns a single user with the password  extracted
func DbUserToUser(dbUsers models.DbUser) models.User {
	var ret models.User
	ret.UserEmail = dbUsers.UserEmail
	ret.ID = dbUsers.ID
	ret.UserName = dbUsers.UserName
	ret.Active = dbUsers.Active
	ret.OnlineStatus = dbUsers.OnlineStatus
	return ret
}
