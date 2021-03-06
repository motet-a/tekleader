package tekleader

import (
	"errors"
	"net/http"
)

// Check if AuthKey is valid
func Auth() error {
	req, err := http.NewRequest("GET", intraURL+AuthKey+"/user", nil)
	if err != nil {
		return err
	}

	res, err := getResRequest(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("Wrong AuthKey - Visit https://intra.epitech.eu/admin/autolog")
	}

	return nil
}
