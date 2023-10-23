package helpers

import (
	"encoding/json"
	"net/http"
)

// BindRequestJSON bind request in JSON format to given data struct
func BindRequestJSON(r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}

	return nil
}