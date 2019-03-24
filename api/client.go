package api

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/agirot/crawler-maxtrain/data"
	"github.com/parnurzeal/gorequest"
	"log"
	"time"
)

//Encode API URL to hide repository in github research
const apiUrlEncode = "aHR0cHM6Ly93d3cub3VpLnNuY2YvYXBpbS9jYWxlbmRhci90cmFpbi92NC8ldi8ldi8ldi8ldi8xMi1IQVBQWV9DQVJELzIvZnI/YWRkaXRpb25hbEZpZWxkcz1ob3Vycw=="

var apiUrl string

func init() {
	url, err := base64.StdEncoding.DecodeString(apiUrlEncode)
	if err != nil {
		log.Panicf(err.Error())
	}

	apiUrl = string(url)
}

func GetCalendar(from, to string) (data.Calendar, error) {
	var calendar data.Calendar

	now := time.Now()
	futur := now.AddDate(0, 0, 31)

	url := fmt.Sprintf(
		apiUrl,
		from,
		to,
		now.Format(data.LayoutDate),
		futur.Format(data.LayoutDate),
	)

	resp, _, errs := gorequest.New().
		Get(url).Timeout(15 * time.Second).EndStruct(&calendar)

	if len(errs) > 0 {
		if resp != nil {
			return calendar, errors.New(fmt.Sprintf("SendAlert failed status code %v", resp.StatusCode))
		}
		return calendar, errs[0]
	} else if resp.StatusCode != 200 {
		return calendar, errors.New(fmt.Sprintf("SendAlert failed status code %v", resp.StatusCode))
	}

	return calendar, nil
}
