package main

import (
	"fmt"
	"regexp"
	"strings"

	ics "github.com/arran4/golang-ical"
)

func filterCalendar(mentions []Mention, groups []Group, options []Option, optionGroups []OptionGroup) ics.Calendar {

	var eventToBanRegex []string
	var filteredCal = *cal

	// Filter course by mention
	for _, mention := range mentions {
		switch mention {
		case "IDL":
			fmt.Println("IDL")
			eventToBanRegex = append(eventToBanRegex, CODE_SDD_AA, CODE_PROBA_POUR_INFORMATIQUE, CODE_CRYPTOGRAPHIE, CODE_PROG_CPP, CODE_PROG_FONCTIONNELLE, CODE_METHODE_NUM_INFORMATIQUE)

		case "I3A":
			fmt.Println("I3A")

		case "FSI":
			fmt.Println("FSI")

		case "GIG":
			fmt.Println("GIG")

		case "IMG":
			fmt.Println("IMD")

		case "SID":
			fmt.Println("SID")

		}
	}

	// Remove unselected groups and add filter
	groupsToBan := []string{"TD1", "TD2", "TD3|TD Gr3", "TD4", "TP1|TP Gr1", "TP2|TP Gr2", "TP3|TP Gr3", "TP4|TP 4"}
	for _, group := range groups {
		switch group {
		case "TD1":
			groupsToBan = removeString(groupsToBan, "TD1")
		case "TD2":
			groupsToBan = removeString(groupsToBan, "TD2")
		case "TD3":
			groupsToBan = removeString(groupsToBan, "TD3|TD Gr3")
		case "TD4":
			groupsToBan = removeString(groupsToBan, "TD4")
		case "TP1":
			groupsToBan = removeString(groupsToBan, "TP1|TP Gr1")
		case "TP2":
			groupsToBan = removeString(groupsToBan, "TP2|TP Gr2")
		case "TP3":
			groupsToBan = removeString(groupsToBan, "TP3|TP Gr3")
		case "TP4":
			groupsToBan = removeString(groupsToBan, "TP4|TP 4")
		}
	}
	eventToBanRegex = append(eventToBanRegex, groupsToBan...)

	// Filter events with options
	for _, option := range options {
		switch option {
		case "cpp":
			eventToBanRegex = removeString(eventToBanRegex, CODE_PROG_CPP)
		case "crypto":
			eventToBanRegex = removeString(eventToBanRegex, CODE_CRYPTOGRAPHIE)
		case "intro-science-donnees":
			eventToBanRegex = removeString(eventToBanRegex, CODE_SDD_AA)
		case "methode-numeriques":
			eventToBanRegex = removeString(eventToBanRegex, CODE_METHODE_NUM_INFORMATIQUE)
		case "prog-fonctionnelle":
			eventToBanRegex = removeString(eventToBanRegex, CODE_PROG_FONCTIONNELLE)
		case "proba":
			eventToBanRegex = removeString(eventToBanRegex, CODE_PROBA_POUR_INFORMATIQUE)
		case "securite-des-apps":
			eventToBanRegex = removeString(eventToBanRegex, CODE_SECURITE_DES_APPS)
		}
	}

	// Filter by option group
	optionGroupToBan := []string{}
	allOptionGroups := []string{"A1", "A2", "A3", "C1", "C2", "F1", "F2"}
	for _, optionGroup := range allOptionGroups {
		found := false
		for _, selectedGroup := range optionGroups {
			if string(selectedGroup) == optionGroup {
				found = true
				break
			}
		}
		if !found {
			optionGroupToBan = append(optionGroupToBan, optionGroup)
		}
	}
	eventToBanRegex = append(eventToBanRegex, optionGroupToBan...)

	// Filter events with the regex created previously
	fmt.Println(strings.Join(eventToBanRegex, "|"))
	for _, value := range filteredCal.Events() {
		if value.GetProperty(ics.ComponentProperty(ics.PropertySummary)) != nil {
			summary := value.GetProperty(ics.ComponentProperty(ics.PropertySummary)).Value
			matched, _ := regexp.MatchString(strings.Join(eventToBanRegex, "|"), summary)
			if matched {
				filteredCal.RemoveEvent(value.Id())
			}
		}
	}

	return filteredCal

}

func removeString(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
