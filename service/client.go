package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

func SendRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil
}
