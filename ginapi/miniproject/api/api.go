package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"najottalim/january-22/miniproject/model"
	"najottalim/january-22/miniproject/storage"
	"net/http"
	"strconv"
)

func main() {

	router := gin.New()

	router.POST("/create", CreateUser)
	router.GET("/all", GetAllUsers)
	router.GET("/getone/:id", GetUserByID)
	router.POST("/update", UpdateUser)
	router.POST("/delete", DeleteUser)
	err := router.Run("localhost:8070")

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

func CreateUser(ctx *gin.Context) {
	//bodyByte, err := io.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	log.Println("Error while reading body", err)
	//	ctx.AbortWithError(http.StatusBadRequest, err)
	//	return
	//}
	// BU ikkinchi usul
	var user *model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Println("Error while json unmarshalling", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := uuid.NewString()
	user.ID = id

	respUser, err := storage.CreateUser(user)
	if err != nil {
		log.Println("Error while creating user error ", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//respBody, err := json.Marshal(respUser)
	//if err != nil {
	//	log.Println("Error while json marshalling error ", err)
	//	ctx.AbortWithError(http.StatusInternalServerError,err)
	//	return
	////}
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//w.Write(respBody)
	ctx.JSONP(http.StatusCreated, respUser)
}

func GetAllUsers(ctx *gin.Context) {
	page := ctx.Request.URL.Query().Get("page")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error while converting page", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit := ctx.Request.URL.Query().Get("limit")

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Error while converting limit", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	users, err := storage.GetAll(intPage, intLimit)
	if err != nil {
		log.Println("Error while getting function not fount", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//respBody, err := json.Marshal(users)
	//if err != nil {
	//	log.Println("Error while json marshalling error ", err)
	//	ctx.AbortWithError(http.StatusInternalServerError, err)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(respBody)

	ctx.JSONP(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	//bodyByte, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Println("Error while reading body", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	var user *model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("Error while json unmarshalling", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	respUsers, err := storage.UpdatedUser([]*model.User{user})
	if err != nil {
		log.Println("Error while updating user", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	//respBody, err := json.Marshal(respUsers)
	//if err != nil {
	//	log.Println("Error while json marshalling error ", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(respBody)
	c.JSONP(http.StatusOK, respUsers)

}

func DeleteUser(c *gin.Context) {
	//bodyByte, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Println("Error while reading body", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	var user *model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("Error while json unmarshalling", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	respUsers, err := storage.DeleteUser([]*model.User{user})
	if err != nil {
		log.Println("Error while updating user", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	//respBody, err := json.Marshal(respUsers)
	//if err != nil {
	//	log.Println("Error while json marshalling error ", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(respBody)
	c.JSONP(http.StatusOK, respUsers)
}
