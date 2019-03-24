package manager

import (
	"fmt"
	"github.com/agirot/crawler-maxtrain/alert"
	"github.com/agirot/crawler-maxtrain/api"
	"github.com/agirot/crawler-maxtrain/data"
	"time"
)

//scheduleFound without mutex (single process)
var scheduleFound = map[string]time.Time{}

func Process(watchDay data.WatchDay) error {
	calendar, err := api.GetCalendar(watchDay.From, watchDay.To)
	if err != nil {
		return err
	}

	for _, day := range calendar {
		schedulesFound, err := performDay(day, watchDay)
		if err != nil {
			return err
		}

		if len(schedulesFound) > 0 {
			err := alert.SendAlert(schedulesFound)
			if err != nil {
				return err
			}

			for _, schedule := range schedulesFound {
				storeSchedule(watchDay, schedule)
			}
		}
	}
	return nil
}

func performDay(schedule data.Schedule, watchWeekDay data.WatchDay) ([]time.Time, error) {
	var match []time.Time
	//We only want date with 0 price value (fast check to see if day contain any TG*V M@X)
	if schedule.Price != 0 {
		return match, nil
	}

	weekDay, err := schedule.GetWeekDay()
	if err != nil {
		return match, err
	}
	//Check if date match with a week day of config
	if watchWeekDay.Day == weekDay.String() {
		//check each schedule to see if match with config
		for _, hours := range schedule.Hours {
			scheduleHour, _ := time.Parse(data.LayoutDateWithTime, fmt.Sprintf("%v %v:00", schedule.Date, hours))
			minConfHour, _ := time.Parse(data.LayoutDateWithTime, fmt.Sprintf("%v %v:00", schedule.Date, watchWeekDay.StartWatch))
			maxConfHour, _ := time.Parse(data.LayoutDateWithTime, fmt.Sprintf("%v %v:00", schedule.Date, watchWeekDay.EndWatch))
			if scheduleHour.After(minConfHour) && scheduleHour.Before(maxConfHour) && !findStoredSchedule(watchWeekDay, scheduleHour) {
				match = append(match, scheduleHour)
			}
		}
	}

	return match, nil
}

func storeSchedule(watchDay data.WatchDay, schedule time.Time) {
	scheduleFound[buildIndex(watchDay, schedule)] = schedule
}

func findStoredSchedule(watchDay data.WatchDay, schedule time.Time) bool {
	_, found := scheduleFound[buildIndex(watchDay, schedule)]
	return found
}

func buildIndex(watchDay data.WatchDay, schedule time.Time) string {
	return fmt.Sprintf("%v_%v_%v_%v", watchDay.From, watchDay.To, watchDay.Day, schedule.UnixNano())
}
