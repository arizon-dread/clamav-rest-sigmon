// api is where controllers are declarated
package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/arizon-dread/clamav-rest-sigmon/internal/service"
	"github.com/arizon-dread/clamav-rest-sigmon/internal/utils"
)

// SignHandler accepts max age either from query string, or falls back to env variable or after that a hard coded value
func SignHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	opts := utils.GetOpts()
	q := r.URL.Query()
	var maxAgeHours int64
	maxAgeHours, err := strconv.ParseInt(opts["MAX_SIGNATURE_AGE_HOURS"], 10, 64)
	if err != nil {
		http.Error(w, `{"error": "Could not get maxAgeHours from query string or from config. This should never occur"}`, http.StatusInternalServerError)
		return
	}
	if q.Get("maxAgeHours") != "" {
		maxAgeH, err := strconv.ParseInt(q.Get("maxAgeHours"), 10, 64)
		if err == nil && maxAgeH != 0 {
			maxAgeHours = maxAgeH
		}
	}
	if maxAgeHours == 0 {
		http.Error(w, `{"error": "Could not get maxAgeHours from query string or from config. This should never occur"}`, http.StatusInternalServerError)
		return
	}
	signatureAge, err := service.CompareSignAge(maxAgeHours)
	if err != nil {
		http.Error(w, err.Error(), 420)
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"message": "Signatures are up-to-date, last check was %d hours ago"}`, signatureAge)))
}
