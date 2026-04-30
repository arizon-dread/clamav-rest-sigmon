package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/arizon-dread/clamav-rest-sigmon/internal/utils"
)

var httpClient *http.Client

func getClient() *http.Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return httpClient
}

// GetClamavSignatureAge returns the number of hours that has passed since the last signature update
func getClamavSignatureAge(opts map[string]string) (int64, int, error) {
	c := getClient()
	res, err := c.Get(fmt.Sprintf("%v/version", opts["CLAMAV_REST_URL"]))
	if err != nil {
		return 0, 500, err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, 500, err
	}
	defer res.Body.Close()

	var v version
	err = json.Unmarshal(b, &v)
	if err != nil {
		return 0, 500, err
	}
	signatureDate, err := time.Parse("Mon Jan 2 15:04:05 2006", v.SignatureDate)
	if err != nil {
		log.Printf("Failed converting timestamp to time.Time, %v", err)
		return 0, 500, err
	}
	now := time.Now().Unix()
	delta := now - signatureDate.Unix()
	deltaHours := delta / 60 / 60
	return deltaHours, 200, nil
}

func CompareSignAge(maxAgeHours int64) (int64, int, error) {
	opts := utils.GetOpts()
	signatureAge, responseCode, err := getClamavSignatureAge(opts)
	if err != nil {
		errf := strings.ReplaceAll(err.Error(), `"`, `'`)
		return signatureAge, responseCode, fmt.Errorf("could get signature age from clamav-rest, %v", errf)
	}
	if signatureAge > maxAgeHours {
		return signatureAge, 420, fmt.Errorf("signatures haven't updated in %d hours", signatureAge)
	}
	return signatureAge, 200, nil
}
