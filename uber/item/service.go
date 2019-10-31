package item

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const hostname = "https://www.ubereats.com/"
const itemApiPath = "api/getMenuItemV1"

func GetItemDetails(body RequestBody) Item {
	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", hostname+itemApiPath, bytes.NewBuffer(requestBody))
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

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		log.Fatalln(err)
	}

	if response.Status != "success" {
		log.Fatalln("Get store request was not successful.")
	}

	return response.Data
}
