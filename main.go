package main

import (
	"fmt"
	"time"
	
	"github.com/arran4/golang-ical@v0.3.2"
)

var weekday string

func init() {
	weekday = time.Now().Weekday().String()
}

func main() {
	fmt.Printf("Today is %s", weekday)
}