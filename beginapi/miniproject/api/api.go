package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"najottalim/january-22/miniproject/model"
	"najottalim/january-22/miniproject/storage"
	"net/http"
	"strconv"
)

func main() {
	//endpoint create user
	http.HandleFunc("/user/create", CreateUser)
	//endpoint get all user
	http.HandleFunc("/user/all", GetAllUsers)
	//updated user
	http.HandleFunc("/user/update", UpdateUser)
	//deleted user
	http.HandleFunc("/user/delete", DeleteUser)
	err := http.ListenAndServe("localhost:8088", nil)
	if err != nil {
		fmt.Println("Error while running server", err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *model.User
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		log.Println("Error while json unmarshalling", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	user.ID = id

	respUser, err := storage.CreateUser(user)
	if err != nil {
		log.Println("Error while creating user error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("Error while json marshalling error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error while converting page", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := r.URL.Query().Get("limit")

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Error while converting limit", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := storage.GetAll(intPage, intLimit)
	if err != nil {
		log.Println("Error while getting function not fount", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(users)
	if err != nil {
		log.Println("Error while json marshalling error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *model.User
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		log.Println("Error while json unmarshalling", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respUsers, err := storage.UpdatedUser([]*model.User{user})
	if err != nil {
		log.Println("Error while updating user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(respUsers)
	if err != nil {
		log.Println("Error while json marshalling error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *model.User
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		log.Println("Error while json unmarshalling", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respUsers, err := storage.DeleteUser([]*model.User{user})
	if err != nil {
		log.Println("Error while updating user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(respUsers)
	if err != nil {
		log.Println("Error while json marshalling error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
