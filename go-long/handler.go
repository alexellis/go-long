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

	if xSleep := r.Header.Get("X-Sleep"); len(xSleep) > 0 {
		sleepDuration, _ = time.ParseDuration(xSleep)
	}

	log.Printf("Start sleep for: %s\n", sleepDuration.String())
	time.Sleep(sleepDuration)
	log.Printf("Sleep done for: %s\n", sleepDuration.String())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Had a nice sleep for: %s", sleepDuration.String())))
}
