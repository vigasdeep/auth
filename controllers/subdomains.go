package controllers

import (
	"auth/models"
	"auth/utils/auth"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//RequestSubdomain function -- request a new subdomain
func RequestSubdomain(w http.ResponseWriter, r *http.Request) {

	// FromContext(r)
	// fmt.Printf("%+v\n", ctx.Value("user"))
	var token, error = auth.VerifyToken(r)
	if error != nil {
		// return nil
	}
	fmt.Printf("%+v\n", token)

	subdomain := &models.Subdomains{}
	json.NewDecoder(r.Body).Decode(subdomain)
	createdSubdomain := db.Create(subdomain)
	var errMessage = createdSubdomain.Error
	// auth.getUserFromJWT(r)

	if createdSubdomain.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdSubdomain)
}

//FetchSubdomains function -- request a new subdomain
func FetchSubdomains(w http.ResponseWriter, r *http.Request) {
	var sd []models.Subdomains
	db.Preload("auths").Find(&sd)

	json.NewEncoder(w).Encode(sd)
}

//DeleteSubdomain function -- request a new subdomain
func DeleteSubdomain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var sd models.Subdomains
	db.First(&sd, id)
	db.Delete(&sd)
	json.NewEncoder(w).Encode("Subdomain deleted")
}

// type ErrorResponse struct {
// 	Err string
// }

// type error interface {
// 	Error() string
// }

// var db = utils.ConnectDB()

// //FetchUser function
// func FetchUsers(w http.ResponseWriter, r *http.Request) {
// 	var users []models.User
// 	db.Preload("auths").Find(&users)

// 	json.NewEncoder(w).Encode(users)
// }

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	user := &models.User{}
// 	params := mux.Vars(r)
// 	var id = params["id"]
// 	db.First(&user, id)
// 	json.NewDecoder(r.Body).Decode(user)
// 	db.Save(&user)
// 	json.NewEncoder(w).Encode(&user)
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var id = params["id"]
// 	var user models.User
// 	db.First(&user, id)
// 	db.Delete(&user)
// 	json.NewEncoder(w).Encode("User deleted")
// }

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var id = params["id"]
// 	var user models.User
// 	db.First(&user, id)
// 	json.NewEncoder(w).Encode(&user)
// }
