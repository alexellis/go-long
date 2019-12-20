package function

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	sleepVal := os.Getenv("handler_wait_duration")
	sleepDuration, _ := time.ParseDuration(sleepVal)
	sleepDurationStr := sleepDuration.String()

	log.Printf("Start sleep for: %s\n", sleepDurationStr)
	time.Sleep(sleepDuration)
	log.Printf("Sleep done for: %s\n", sleepDurationStr)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Had a nice sleep for: %s", sleepDuration.String())))
}
