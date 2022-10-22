package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPhone(id string) ([]byte, int, error) {
	url := fmt.Sprintf("http://localhost:8008/phone/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get phone from db: %s", err)
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}
