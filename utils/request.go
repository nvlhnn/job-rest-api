package utils

import (
	"dansmultipro/recruitment/dto"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func GetResponses(url string) []dto.JobResponse{
	var responses []dto.JobResponse

	var myClient = &http.Client{Timeout: 10 * time.Second}

	resp, err := myClient.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Println("migrate failed on get")
		log.Println(err)
		return nil
	}
	
	err = json.NewDecoder(resp.Body).Decode(&responses)

	return responses
}

func GetResponse(url string, id string) dto.JobResponse{
	var response dto.JobResponse

	var myClient = &http.Client{Timeout: 10 * time.Second}

	resp, err := myClient.Get(url+ "/" +id)
	defer resp.Body.Close()

	if err != nil {
		log.Println("migrate failed on get")
		log.Println(err)
		return response
	}
	
	err = json.NewDecoder(resp.Body).Decode(&response)

	return response
}