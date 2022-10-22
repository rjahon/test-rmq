package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetPhone(id int) ([]byte, int, error) {
	url := fmt.Sprintf("https://127.0.0.1:8008/phone/%s", strconv.Itoa(id))
	resp, err := http.Get(url)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

func ParseId(args []string) int {
	var (
		id  int
		err error
	)

	if len(args) > 1 || os.Args[2] == "" {
		return 1
	} else {
		id, err = strconv.Atoi(args[0])
		if err != nil {
			log.Printf("%s: %s", "Failed to parse id from args", err)
			return 1
		}
	}
	return id
}
