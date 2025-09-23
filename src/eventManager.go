package main

import (
	"time"

	ics "github.com/arran4/golang-ical"
)

func StartPeriodicFetching() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	
	// Fetch immediately
	cal, _ = ics.ParseCalendarFromUrl(generateURL())

	// Then fetch every hour
	for range ticker.C {
		cal, _ = ics.ParseCalendarFromUrl(generateURL())
	}
}

func generateURL() string {
	return "https://ade-web-consult.univ-amu.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?projectId=8&resources=1819&calType=ical&firstDate=" + time.Now().Format("2006-01-02") + "&lastDate=2026-09-19"
}
