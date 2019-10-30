package store

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const hostname = "https://www.ubereats.com/"
const storeApiPath = "api/getStoreV1"

func GetStoreDetails(uuid string) Data {
	requestBody, err := json.Marshal(map[string]interface{}{
		"storeUuid":     uuid,
		"sfNuggetCount": 0,
	})
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", hostname+storeApiPath, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-csrf-token", "x")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	if response.Status != "success" {
		log.Fatalln("Get store request was not successful.")
	}

	return response.Data
}
