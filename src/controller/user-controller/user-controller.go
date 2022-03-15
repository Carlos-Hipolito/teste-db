package UserController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"teste-db/src/database"
	"teste-db/src/models"

	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDatabase()
	var users models.Users
	params := mux.Vars(r)
	db.First(&users, params["id"])
	x := json.NewDecoder(r.Body).Decode(&users)
	db.Save(&users)
	json.NewEncoder(w).Encode(users)
	fmt.Println("Usuario com a ID ", x, "atualizado com sucesso.")
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDatabase()
	var users models.Users
	params := mux.Vars(r)
	db.Find(&users, params["id"])
	json.NewEncoder(w).Encode(users)
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDatabase()
	var users models.Users
	db.Find(&users)
	allusers, err := json.Marshal(&users)
	if err != nil {
		fmt.Println("error when try to Marshal allusers ")
	}
	w.Write(allusers)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var newuser models.User
	_ = json.NewDecoder(r.Body).Decode(&newuser)
	db.Create(&newuser)
	fmt.Fprint(w, "Usuario cadastrado com sucesso com o ID: ", newuser.ID)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var users models.Users
	var idParam string = mux.Vars(r)["id"]
	userID, err := strconv.Atoi(idParam)
	fmt.Println(userID)
	if err != nil {
		fmt.Println("error when try to convert idParam to int")
	}
	db.Delete(&users, userID)
	fmt.Fprint(w, "Usuario com a id ", (userID), " deletado com sucesso.")
}

func Teste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Route TESTE is working")
}

func ExecuteAll(r *mux.Router) {

	r.HandleFunc("/users", UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", DeleteUserByID).Methods("DELETE")
	r.HandleFunc("/users/{id}", FindUserByID).Methods("GET")
	r.HandleFunc("/teste", Teste).Methods("GET")
	r.HandleFunc("/users", FindAll).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")

}
