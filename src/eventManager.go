package main

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

func StartPeriodicFetching() {
	var err error

	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	// Fetch immediately
	fmt.Println("Fetching calendar...")
	cal, err = ics.ParseCalendarFromUrl(generateURL())
	if err != nil {
		fmt.Println("Error fetching calendar:", err)
		return
	}

	// Then fetch every hour
	for range ticker.C {
		fmt.Println("Fetching calendar...")
		cal, err = ics.ParseCalendarFromUrl(generateURL())
		if err != nil {
			fmt.Println("Error fetching calendar:", err)
			continue
		}
	}
}

func generateURL() string {
	return "https://ade-web-consult.univ-amu.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?projectId=8&resources=1819&calType=ical&firstDate=" + time.Now().Format("2006-01-02") + "&lastDate=2026-09-19"
}
