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
		apiErr : &utils.ApplicationErro{
			Message: "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue, _ := json.Marshal()
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
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