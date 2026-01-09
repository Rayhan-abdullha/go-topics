package handler

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/repo"
	"server/utils"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	usr := repo.User{}
	json.NewDecoder(r.Body).Decode(&usr)
	if usr.Name == "" || usr.Email == "" {
		utils.ErrorData(w, "Bad Request Data", http.StatusBadRequest)
		return
	}
	findUser := repo.FindUserByEmail(usr.Email)
	if findUser != nil {
		utils.ErrorData(w, "User already exists", http.StatusConflict)
		return
	}
	usr.ID = len(repo.UserList) + 1
	repo.UserList = append(repo.UserList, usr)

	payload := utils.Payload{
		Sub:   usr.ID,
		Name:  usr.Name,
		Email: usr.Email,
	}
	cnf := config.GetConfig()
	token, _ := utils.CreateJWT(payload, cnf.SecretJWT)
	utils.SendData(w, token, 201)
}
