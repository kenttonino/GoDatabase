package sqlite

import (
	"GoDatabase/src/services"
	"GoDatabase/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func routeError(err string, w http.ResponseWriter) {
	log.Println(utils.TextRed(err))
	errResponse := &utils.HTTPResponse{Message: err, Status: http.StatusBadRequest}
	errResponseJson, _ := json.Marshal(errResponse)
	fmt.Fprint(w, string(errResponseJson))
}

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	route := utils.SQLReadyRoute + " "

	db, err := services.CreateSQLInstance()

	if err != nil {
		routeError(err.Error(), w)
		return
	}
	defer db.Close()

	response := &utils.HTTPResponse{Message: "SQLite is ready", Status: http.StatusOK}
	responseJson, _ := json.Marshal(response)
	log.Print(utils.TextGreen(route + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}

func CreateUserTableHandler(w http.ResponseWriter, r *http.Request) {
	route := utils.SQLCreateUserTableRoute + " "
	err := services.CreateSQLUserTable()

	if err != nil {
		routeError(err.Error(), w)
		return
	}

	response := &utils.HTTPResponse{Message: "Successfully created the user table.", Status: http.StatusOK}
	responseJson, _ := json.Marshal(response)
	log.Print(utils.TextGreen(route + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}
