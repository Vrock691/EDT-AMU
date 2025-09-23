package main

import (
	"fmt"
	"regexp"
	"strings"

	ics "github.com/arran4/golang-ical"
)

func filterCalendar(mentions []Mention, groups []Group, options []Option, optionGroups []OptionGroup) ics.Calendar {

	var eventToRemoveRegex []string
	filteredCal := *cal

	for _, mention := range mentions {
		if codes, exists := mentionToCodesMap[string(mention)]; exists {
			eventToRemoveRegex = append(eventToRemoveRegex, codes...)
		}
	}

	// Remove unselected groups and add filter
	groupsToBan := make([]string, 0, len(groupToRegexMap))
	for _, pattern := range groupToRegexMap {
		groupsToBan = append(groupsToBan, pattern)
	}

	for _, group := range groups {
		if pattern, exists := groupToRegexMap[group]; exists {
			groupsToBan = removeStringFromList(groupsToBan, pattern)
		}
	}
	eventToRemoveRegex = append(eventToRemoveRegex, groupsToBan...)

	// Filter events with options
	for _, option := range options {
		switch option {
		case "cpp":
			eventToRemoveRegex = removeStringFromList(eventToRemoveRegex, CODE_PROG_CPP)
		case "crypto":
			eventToRemoveRegex = removeStringFromList(eventToRemoveRegex, CODE_CRYPTOGRAPHIE)
		case "intro-science-donnees":
			eventToRemoveRegex = removeStringFromList(eventToRemoveRegex, CODE_SDD_AA)
		case "methode-numeriques":
			eventToRemoveRegex = removeStringFromList(eventToRemoveRegex, CODE_METHODE_NUM_INFORMATIQUE)
		case "prog-fonctionnelle":
			eventToRemoveRegex = removeStringFromList(eventToRemoveRegex, CODE_PROG_FONCTIONNELLE)
		case "proba":
			eventToRemoveRegex = removeStringFromList(eventToRemoveRegex, CODE_PROBA_POUR_INFORMATIQUE)
		case "securite-des-apps":
			eventToRemoveRegex = removeStringFromList(eventToRemoveRegex, CODE_SECURITE_DES_APPS)
		}
	}

	// Filter by option group
	allOptionGroups := make([]string, 0, len(optionGroupNumberToRegexMap))
	for _, pattern := range optionGroupNumberToRegexMap {
		allOptionGroups = append(allOptionGroups, pattern)
	}
	for _, optionGroup := range optionGroups {
		if pattern, exists := optionGroupNumberToRegexMap[string(optionGroup)]; exists {
			allOptionGroups = removeStringFromList(allOptionGroups, pattern)
		}
	}
	eventToRemoveRegex = append(eventToRemoveRegex, allOptionGroups...)

	// Filter events with the regex created previously
	fmt.Println(strings.Join(eventToRemoveRegex, "|"))
	for _, value := range filteredCal.Events() {
		if value.GetProperty(ics.ComponentProperty(ics.PropertySummary)) != nil {
			summary := value.GetProperty(ics.ComponentProperty(ics.PropertySummary)).Value
			matched, _ := regexp.MatchString(strings.Join(eventToRemoveRegex, "|"), summary)
			if matched {
				filteredCal.RemoveEvent(value.Id())
			}
		}
	}

	return filteredCal

}

func removeStringFromList(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
