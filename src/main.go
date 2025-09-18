package main

import (
	"fmt"
	"net/http"

	ics "github.com/arran4/golang-ical"
)

var cal *ics.Calendar

func main() {

	// Start the event fetcher in background
	go StartPeriodicFetching()

	// Define the front route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/html/index.html")
	})

	// Define the ics route
	http.HandleFunc("/calendar.ics", func(w http.ResponseWriter, r *http.Request) {

		// TODO: Add filtering

		fmt.Println("Obtention du calendrier")
		w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
		w.Header().Set("Content-Disposition", "attachment; filename=calendar.ics")
		if cal != nil {
			fmt.Fprint(w, cal.Serialize())
		}
	})

	// Start the http server
	fmt.Println("Starting HTTP server on :8080")
	http.ListenAndServe(":8080", nil)
}
