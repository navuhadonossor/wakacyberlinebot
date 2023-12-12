package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const userInfoUrl = "/api/compat/wakatime/v1/users/current"
const userStatUrl = "/api/compat/wakatime/v1/users/current/stats/"

func getWakaUserInfo(apiToken string) wakatimeUserResponse {
	var jsonResponse wakatimeUserResponse
	bytes := wakaRequest(apiToken, os.Getenv("WAKATIME_HOST")+userInfoUrl)
	_ = json.Unmarshal(bytes, &jsonResponse)
	return jsonResponse
}

func getUserStatInfo(apiToken string, period string) wakatimeStatResponse {
	var jsonResponse wakatimeStatResponse
	bytes := wakaRequest(apiToken, os.Getenv("WAKATIME_HOST")+userStatUrl+period)
	_ = json.Unmarshal(bytes, &jsonResponse)
	return jsonResponse
}

func wakaRequest(apiToken string, url string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Failed to construct request!")
	}
	request.Header.Add("Authorization", "Basic "+apiToken)
	response, err := client.Do(request)
	if err != nil {
		log.Println("Response failed: " + err.Error())
	}
	bytes, _ := io.ReadAll(response.Body)
	return bytes
}
