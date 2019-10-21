package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/learning-dev/golang-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userId, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		// just return the Bad Request to the client.
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number"))
		return
	}
	user, err := services.GetUser(userId)
	if err != nil{
		// Handle the err and return to the client
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}