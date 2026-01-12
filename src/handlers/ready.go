package handlers

import (
	"GoDatabase/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	route := utils.ReadyRoute + " "
	response := &utils.HTTPResponse{Message: "Server is ready.", Status: 200}
	responseJson, _ := json.Marshal(response)

	log.Print(utils.TextGreen(route + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}
