package main

import (
	"time"

	ics "github.com/arran4/golang-ical"
)

var adeURL string = "https://ade-web-consult.univ-amu.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?projectId=8&resources=1819&calType=ical&firstDate=2025-09-15&lastDate=2026-09-19"
var eventManagerError error

func StartPeriodicFetching() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	// Fetch immediately
	cal, eventManagerError = ics.ParseCalendarFromUrl(adeURL)

	// Then fetch every hour
	for range ticker.C {
		cal, eventManagerError = ics.ParseCalendarFromUrl(adeURL)
	}
}
