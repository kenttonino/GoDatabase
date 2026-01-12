package handlers

import (
	"GoDatabase/src/services"
	"GoDatabase/src/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func RedisReadyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	route := utils.RedisReadyRoute + " "
	redisClient := services.NewRedisClient()
	_, err := redisClient.Ping(ctx).Result()

	if err != nil {
		log.Print(utils.TextRed(route + strconv.Itoa(500)))
		errorResponse := &utils.HTTPResponse{Message: "Failed redis connection.", Status: 500}
		errorResponseJson, _ := json.Marshal(errorResponse)
		fmt.Fprint(w, string(errorResponseJson))
		return
	}

	response := &utils.HTTPResponse{Message: "Redis is ready.", Status: 200}
	responseJson, _ := json.Marshal(response)
	log.Print(utils.TextGreen(route + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}
