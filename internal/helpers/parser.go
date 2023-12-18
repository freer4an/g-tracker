package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseAPI(url string, n any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("Response body closing error: %v", err)
		}
	}()

	return json.NewDecoder(resp.Body).Decode(&n)
}
