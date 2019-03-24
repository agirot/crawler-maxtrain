package data

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const LayoutDate = "2006-01-02"
const LayoutDateWithTime = "2006-01-02 15:04:05"

var daysOfWeek = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
}

//WatchDayConfig is configuration file struct
type WatchDaysConfig []WatchDay

type WatchDay struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Day        string `json:"day"`
	StartWatch string `json:"start_watch"`
	EndWatch   string `json:"end_watch"`
}

//Calendar is struct of API output result
type Calendar []Schedule

// Var of flag start
var SmsKeyArg *string
var SmsUserArg *string
var ConfigPath string

var Config WatchDaysConfig

//Schedule is API data struct
type Schedule struct {
	Date  string   `json:"date"`
	Price int      `json:"price"`
	Hours []string `json:"hours"`
}

//Convert string date to Weekday type
func (d Schedule) GetWeekDay() (time.Weekday, error) {
	t, err := time.Parse(LayoutDate, d.Date)
	return t.Weekday(), err
}

func CheckDayExist(day string) error {
	if _, ok := daysOfWeek[strings.Title(day)]; !ok {
		return errors.New(fmt.Sprintf("%v is a invalid day of week !", day))
	}
	return nil
}

func GetSmsAuth() (string, string) {
	login := ""
	if SmsKeyArg != nil {
		login = *SmsUserArg
	}

	password := ""
	if SmsKeyArg != nil {
		password = *SmsKeyArg
	}

	return login, password
}


