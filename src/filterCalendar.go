package main

import (
	"fmt"
	"regexp"
	"strings"

	ics "github.com/arran4/golang-ical"
)

/*
Filters the calendar events based on the selected mentions, groups, options, and option groups
*/
func filterCalendar(mentions []Mention, groups []Group, options []Option, optionGroups []OptionGroup) *ics.Calendar {

	// Create a regex expression to remove undesired events
	var eventToExcludeRegex []string

	// Copy the source calendar
	filteredCal := ics.NewCalendar()
	filteredCal.SetName("M1 Informatique")

	// Add to regex every events that are not in the selected mentions
	for _, mention := range mentions {
		if codes, exists := mentionToCodesMap[string(mention)]; exists {
			eventToExcludeRegex = append(eventToExcludeRegex, codes...)
		}
	}

	// Add to regex every events that are not in the selected groups
	groupsToExclude := make([]string, 0, len(groupToRegexMap))
	for _, pattern := range groupToRegexMap {
		groupsToExclude = append(groupsToExclude, pattern)
	}

	for _, group := range groups {
		if pattern, exists := groupToRegexMap[group]; exists {
			groupsToExclude = removeStringFromList(groupsToExclude, pattern)
		}
	}
	eventToExcludeRegex = append(eventToExcludeRegex, groupsToExclude...)

	// Add to regex every events that are not in the selected options
	// As options are already part of others mentions, they're already in the regex, so we just have to remove them from it if they're selected
	for _, option := range options {
		switch option {
		case "cpp":
			eventToExcludeRegex = removeStringFromList(eventToExcludeRegex, CODE_PROG_CPP)
		case "crypto":
			eventToExcludeRegex = removeStringFromList(eventToExcludeRegex, CODE_CRYPTOGRAPHIE)
		case "intro-science-donnees":
			eventToExcludeRegex = removeStringFromList(eventToExcludeRegex, CODE_SDD_AA)
		case "methode-numeriques":
			eventToExcludeRegex = removeStringFromList(eventToExcludeRegex, CODE_METHODE_NUM_INFORMATIQUE)
		case "prog-fonctionnelle":
			eventToExcludeRegex = removeStringFromList(eventToExcludeRegex, CODE_PROG_FONCTIONNELLE)
		case "proba":
			eventToExcludeRegex = removeStringFromList(eventToExcludeRegex, CODE_PROBA_POUR_INFORMATIQUE)
		case "securite-des-apps":
			eventToExcludeRegex = removeStringFromList(eventToExcludeRegex, CODE_SECURITE_DES_APPS)
		}
	}

	// Add to regex every events that are not in the selected option groups
	// We take all the regex of option groups, and remove the selected ones
	// Then we add the remaining ones to the regex to remove
	allOptionGroups := make([]string, 0, len(optionGroupNumberToRegexMap))
	for _, pattern := range optionGroupNumberToRegexMap {
		allOptionGroups = append(allOptionGroups, pattern)
	}
	for _, optionGroup := range optionGroups {
		if pattern, exists := optionGroupNumberToRegexMap[string(optionGroup)]; exists {
			allOptionGroups = removeStringFromList(allOptionGroups, pattern)
		}
	}
	eventToExcludeRegex = append(eventToExcludeRegex, allOptionGroups...)

	// Filter events with the regex created previously
	fmt.Println(strings.Join(eventToExcludeRegex, "|"))
	for _, value := range cal.Events() {
		if value.GetProperty(ics.ComponentProperty(ics.PropertySummary)) != nil {
			summary := value.GetProperty(ics.ComponentProperty(ics.PropertySummary)).Value
			matched, _ := regexp.MatchString(strings.Join(eventToExcludeRegex, "|"), summary)
			// Add event to the filtered calendar if it's excluded by the regex
			if !matched {
				filteredCal.AddEvent(value.Id())
			}
		}
	}

	// Return the filtered calendar
	return filteredCal
}

/*
Removes a string from a list of strings
*/
func removeStringFromList(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
