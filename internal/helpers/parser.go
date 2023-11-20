package helpers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func ParseAPI(url string, n any) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Parse error: %w", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			slog.Warn("Parse error: %w", err)
		}
	}()

	return json.NewDecoder(resp.Body).Decode(&n)
}
