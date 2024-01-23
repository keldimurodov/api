package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"najottalim/january-22/miniproject/model"
	"najottalim/january-22/miniproject/storage"
	"net/http"
	"strconv"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/create", CreateUser).Methods("POST")
	router.HandleFunc("/all", GetAllUsers).Methods("GET")
	//router.HandleFunc("/getone/:id", GetUserByID).Methods("GET")
	router.HandleFunc("/update", UpdateUser).Methods("POST")
	router.HandleFunc("/delete", DeleteUser).Methods("POST")
	err := http.ListenAndServe("localhost:8070", router)

	if err != nil {
		log.Fatal(err)
	}
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := storage.GetUser(userID)
	if err != nil {
		log.Println("Error while getting user", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSONP(http.StatusOK, user)
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
