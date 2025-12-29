package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Mboukhal/FactoryBase/core/roles"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenExpiry      = time.Hour * 72
)

type UserInfo struct {
	ID      		string 				`json:"id"`
	Email   		string 				`json:"email"`
	Name    		string 				`json:"name"`
	Role			roles.ServiceRole	`json:"role"`
	OrganizationID 	string				`json:"organization_id"`
	Picture 		string 				`json:"picture"`
}



func CreateJWT(userInfo UserInfo) (string, error) {

	claims := jwt.MapClaims{
		"user": UserInfo{
			ID:      userInfo.ID,
			Email:   userInfo.Email,
			Name:    userInfo.Name,
			Role:    userInfo.Role,
			OrganizationID: userInfo.OrganizationID,
			Picture: userInfo.Picture,
		},
		"exp": TokenExpiry,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error signing JWT: %s", err.Error())
	}

	return tokenString, nil
}


func GetUserFromDB(email string) (roles.ServiceRole, error) {
	
	
	if email == "" || !strings.Contains(email, "@") {
		return roles.RoleUnknown, fmt.Errorf("invalid email address")
	}

	// var userInfo auth.UserInfo
	// TODO: get user role and name from DB
	// err :=roles.IsValidRole(userInfo.Role)

	err := roles.IsValidRole(roles.RoleAdmin)
	if err != nil {
		return "", err
	} else {
		return roles.RoleAdmin, nil
	}

	// return roles.RoleUnknown, nil
}