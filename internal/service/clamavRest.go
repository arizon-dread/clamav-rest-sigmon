package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var httpClient *http.Client

func getClient() *http.Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return httpClient
}

// GetClamavSignatureAge returns the number of hours that has passed since the last signature update
func GetClamavSignatureAge(opts map[string]string) (int64, error) {
	c := getClient()
	res, err := c.Get(fmt.Sprintf("%v/version", opts["CLAMAV_REST_URL"]))
	if err != nil {
		return 0, err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var v version
	err = json.Unmarshal(b, &v)
	if err != nil {
		return 0, err
	}
	signatureDate, err := time.Parse("Mon Jan 2 15:04:05 2006", v.SignatureDate)
	if err != nil {
		log.Printf("Failed converting timestamp to time.Time, %v", err)
	}
	now := time.Now().Unix()
	delta := now - signatureDate.Unix()
	deltaHours := delta / 60 / 60
	return deltaHours, nil
}
