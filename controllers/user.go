package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Teyik0/go-test/database"
	"github.com/Teyik0/go-test/helpers"
	"github.com/Teyik0/go-test/prisma/db"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient
	allUsers, err := pClient.Client.User.FindMany().Exec(pClient.Context)
	if err != nil {
		fmt.Println("Cannot fetch users")
		return

	}
	usersMap := make(map[string]interface{})
	usersMap["users"] = allUsers

	// This should be handled in a helper funtion
	// But I'm leaving it here to make this easier
	out, err := json.MarshalIndent(usersMap, "", "\t")
	if err != nil {
		fmt.Println("Error Creating JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		fmt.Println("Error writing to response")
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userResp db.UserModel
	err := json.NewDecoder(r.Body).Decode(&userResp)
	if err != nil {
		fmt.Println("Cannot decode user")
		return
	}
	pClient := database.PClient
	user, err := pClient.Client.User.CreateOne(
		db.User.Email.Set(userResp.Email),
		db.User.Password.Set(userResp.Password), // You should hash password !!!
		db.User.Firstname.Set(userResp.Firstname),
		db.User.Lastname.Set(userResp.Lastname),
	).Exec(pClient.Context)
	if err != nil {
		fmt.Println("Cannot create user")
		return
	}
	err = helpers.WriteJSON(w, http.StatusOK, user)
	if err != nil {
		fmt.Println("Cannot form response")
		return
	}
}
