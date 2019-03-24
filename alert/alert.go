package alert

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/agirot/crawler-maxtrain/data"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"text/template"
	"time"
)

const layoutTimeSms = "2006-01-02 15:04:05"

//SendAlert Send alert with free mobile API (my gsm provider)
func SendAlert(schedules []time.Time) error {
	text, err := renderSms(schedules)
	if err != nil {
		return err
	}

	req := gorequest.New()
	login, key := data.GetSmsAuth()
	resp, _, errs := req.Post("https://smsapi.free-mobile.fr/sendmsg").
		Timeout(15 * time.Second).
		Query(fmt.Sprintf("user=%v", login)).
		Query(fmt.Sprintf("pass=%v", key)).
		Query(fmt.Sprintf("msg=%v", text)).
		End()

	if len(errs) > 0 {
		if resp != nil {
			return errors.New(fmt.Sprintf("SendAlert failed status code %v", resp.StatusCode))
		}
		return errs[0]
	} else if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("SendAlert failed status code %v", resp.StatusCode))
	}

	println("Alert send", text)
	return nil
}

func renderSms(schedules []time.Time) (string, error) {
	funcMap := template.FuncMap{
		"formatTime": formatTime,
	}

	var writer bytes.Buffer
	text := `ALERT BOOK NOW !{{ range $schedule := . }}
	{{ formatTime $schedule }}{{end}}`
	t := template.Must(template.New("config").Funcs(funcMap).Parse(text))
	err := t.Execute(&writer, schedules)

	return url.PathEscape(writer.String()), err
}

func formatTime(time time.Time) string {
	return fmt.Sprintf("%v: %v", time.Weekday().String(), time.Format(layoutTimeSms))
}
