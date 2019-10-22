package services

import (
	"github.com/learning-dev/golang-microservices/mvc/domain"
)

func GetUser(userId int64) (*domain.User, *ApplicationError){
	return domain.GetUser(userId)

}
