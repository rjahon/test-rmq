package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPhone(id string) ([]byte, int, error) {
	// i := strconv.Itoa(id)
	url := fmt.Sprintf("http://localhost:8008/phone/%s", id)
	log.Println("url: " + url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%s: %s", "Failed to get phone from db", err)
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}
